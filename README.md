translation
=========================
一个英中翻译命令行工具，接受标准输入（Stdin）的文本，输出翻译文本（Println）。


使用
=========================
不加任何参数为使用 google API 进行翻译, 但需要科学上午。

如果需要使用彩云小译api, 需要在工具的后面加上 ` --p=cyxy --cyxy-token=3975l6lr5pcbvidl6jl2` 参数
* `--p` 为指定翻译服务API: google、cyxy(彩云小译). 默认：google
* `--cyxy-token` 为彩云小译的访问令牌(设定 --p=cyxy 时必须存在)，可以使用 `3975l6lr5pcbvidl6jl2` 作为测试 Token，但官方不保证该 Token 的可用性，所以如果要持续使用，还请申请正式 Token。
    * [彩云小译API](https://fanyi.caiyunapp.com/#/api)
    * 每月翻译100万字之内免费，个人使用够了

### 与 Keyboard Maestro 配合使用
配置参考
![](https://raw.githubusercontent.com/Herbertzz/imgs/master/translation-readme-google-km.png)

操作流程: 选中需要翻译的英文文本 -> 按下设置的快捷键 -> 翻译完成后会自动显示

### 与 Hammerspoon 配合使用
操作流程和 Keyboard Maestro 一样
```lua
-- 选中翻译
hs.hotkey.bind(HyperKey, "T", function()
	-- 配置项
	local toolPath = "/Users/herbertzz/coding/golang/bin/translation"
	local provider = "cyxy"
	local cyxyToken = "3975l6lr5pcbvidl6jl2"
	-- 调用系统的拷贝热键
	hs.eventtap.keyStroke({"cmd"}, "c")
	-- 延迟 1 秒执行
	hs.timer.delayed.new(1, function()
		-- 检查系统剪贴板最新一条记录是否是文本类型
		local isPlainText = false
		local contentType = hs.pasteboard.contentTypes()
		for index, uti in ipairs(contentType) do
			if string.find(uti, 'plain-text', 1, true) ~= nil then
				isPlainText = true
				break
			end
		end
		-- 非文本类型的话，提示错误并中止
		if isPlainText == false then
			hs.alert.show("错误提示: 选中内容非文本", {strokeColor={white=0, alpha=0.75}, textColor={red=1}})
			return false
		end
		-- 将拷贝内容作为标准输入传递给翻译工具，显示标准输出
		local toolParam = {}
		toolParam[1] = "-p"
		toolParam[2] = provider
		if provider == "cyxy" then
			toolParam[3] = "-cyxy-token"
			toolParam[4] = cyxyToken
		end
		task = hs.task.new(toolPath, function(exitCode, stdOut, stdErr)
			-- 调用AppleScript的对话框
			local appleScriptCode = string.format('display dialog "%s" with title "翻译" buttons {"OK"} default button 1', stdOut)
			hs.osascript.applescript(appleScriptCode)
		end, toolParam)
		task:setInput(hs.pasteboard.getContents())
		task:start()
	end):start()
end)
```

编译
=========================