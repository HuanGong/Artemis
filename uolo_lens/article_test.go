package uolo_lens

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"testing"
)

func TestPostHandler_NewPost(t *testing.T) {
	/*
		GET https://mercury.postlight.com/parser?url=https://trackchanges.postlight.com/building-awesome-cms-f034344d8ed
		Content-Type: application/json
		x-api-key: vy0oEGCb4MUl2S4CCwH5UYZRRMWkpS96BEQiE4Pg*/

	client := &http.Client{}
	//生成要访问的url
	url := "https://mercury.postlight.com/parser?url=https://trackchanges.postlight.com/building-awesome-cms-f034344d8edhttp://www.xinhuanet.com/gangao/2018-05/03/c_129864043.htm"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Content-Type", "application/json")
	reqest.Header.Add("x-api-key", "vy0oEGCb4MUl2S4CCwH5UYZRRMWkpS96BEQiE4Pg")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println("body error:", readErr)
	}
	fmt.Println("body:", string(body))

	defer response.Body.Close()
}

func TestPostHandler_ArticleNew(t *testing.T) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	//cmd := exec.Command("clean-mark", "http://tech.ifeng.com/a/20180504/44980993_0.shtml")
	cmd := exec.Command("clean-mark", "http://tech.ifeng.com/a/20180504/44980993_0.shtml")

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	if err != nil {
		fmt.Println("error cmd.run:", err.Error())
	}
	//实时循环读取输出流中的一行内容
	metaStart := false
	for {
		line, err2 := out.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			fmt.Println("error:", err2.Error())
			break
		}

		if strings.HasPrefix(line, "---") {
			metaStart = !metaStart
			if metaStart == false {
				break
			}
			continue
		}

		if false == metaStart {
			continue
		}
		fmt.Println("line:", line)
		Infos := strings.Split(line, ":")
		if len(Infos) < 2 {
			continue
		}
		switch Infos[0] {
		case "link":
			fmt.Println("Link:", Infos[1])
		case "title":
			fmt.Println("Title:", Infos[1])
		case "description":
			fmt.Println("Description:", Infos[1])
		case "keywords":
			fmt.Println("keywords:", Infos[1])
		case "author":
			fmt.Println("author:", Infos[1])
		case "date":
			fmt.Println("date:", strings.Join(Infos[1:], ":"))
		case "publisher":
			fmt.Println("publisher:", Infos[1])
		default:
			continue
		}
	}

	//cmd := exec.Command("clean-mark", "http://tech.ifeng.com/a/20180504/44980993_0.shtml")
	return
}
