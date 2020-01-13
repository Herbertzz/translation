translation
=========================
一个英中翻译命令行工具，接受标准输入（Stdin）的文本，输出翻译文本（Println）。


使用
=========================
不加任何参数为使用 google API 进行翻译, 但需要科学上午。

如果需要使用彩云小译api, 需要在工具的后面加上 ` --p=cyxy --cyxy-token=of5v2usxqqi4cfv2x04m` 参数
* `--p` 为指定翻译服务API: google、cyxy(彩云小译). 默认：google
* `--cyxy-token` 为彩云小译的访问令牌(设定 --p=cyxy 时必须存在)，可以使用 `3975l6lr5pcbvidl6jl2` 作为测试 Token，但官方不保证该 Token 的可用性，所以如果要持续使用，还请申请正式 Token。
    * [彩云小译API](https://fanyi.caiyunapp.com/#/api)
    * 每月翻译100万字之内免费，个人使用够了

### 与 Keyboard Maestro 配合使用
配置参考
![](https://raw.githubusercontent.com/Herbertzz/imgs/master/translation-readme-google-km.png)

操作流程: 选中需要翻译的英文文本 -> 按下设置的快捷键 -> 翻译完成后会自动显示


编译
=========================