package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// 获取命令行参数
	provider := flag.String("p", "google", "服务商提供者: 彩云小译(cyxy), google")
	cyxyToken := flag.String("cyxy-token", "", "彩云小译 API Token")
	flag.Parse()
	// 获取标准输入
	var mtext []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		mtext = append(mtext, sc.Text())
	}

	switch *provider {
	case "google":
		for _, v := range mtext {
			google(v)
		}
		break
	case "cyxy":
		xiaoyi(mtext, *cyxyToken)
		break
	default:
		fmt.Println("暂不支持的服务提供者")
	}
}

// google 翻译
type translation [][][]string
func google(str string) {
	// 参数处理
	apiUrl := "http://translate.google.com/translate_a/single?client=at&sl=en&tl=zh-CN&dt=t&q="
	request := apiUrl + str
	u, _ := url.Parse(request)
	q := u.Query()
	u.RawQuery = q.Encode()

	// 请求api
	response, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 解析请求结构
	translation := translation{}
	json.Unmarshal(body, &translation)
	for _, v := range translation[0] {
		fmt.Println(v[0])
	}
}

// 彩云小译
type XiaoyiRequestBody struct {
	Source    []string `json:"source"`
	TransType string   `json:"trans_type"`
	Replaced  bool     `json:"replaced"`
	Media     string   `json:"media"`
	RequestId string   `json:"request_id"`
}
type XiaoyiResponseBody struct {
	Confidence float32  `json:"confidence"`
	Rc         int      `json:"rc"`
	SrcTgt     []string `json:"src_tgt"`
	Target     []string `json:"target"`
}

func xiaoyi(mtext []string, token string) {
	apiUrl := "http://api.interpreter.caiyunai.com/v1/translator"

	body := XiaoyiRequestBody{
		Source:    mtext,
		TransType: "en2zh",
		Replaced:  true,
		Media:     "text",
		RequestId: "demo",
	}
	reqBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(string(reqBody)))
	if err != nil {
		log.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var rb XiaoyiResponseBody
	err = json.Unmarshal(b, &rb)
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range rb.Target {
		fmt.Println(v)
	}
}
