package uolo_lens

import (
	"encoding/base64"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestPostHandler_ArticleNew(t *testing.T) {
	apiUrl := "http://localhost:3006/lens/article/new"
	form := url.Values{}

	form["title"] = []string{"Hello Kitty Nice To meet you"}
	form["tag"] = []string{"C++"}
	form["author"] = []string{"gonghuan"}
	form["mime"] = []string{"md"}
	form["origin"] = []string{"http://www.baidu.com/news/abc"}
	form["summary"] = []string{"I Have a dream, but life only once in my all life time, so tell myself that 'you only live once'"}
	form["content"] = []string{`# Remember, You Only Live Once
		- do the things you wanna do
		- to be the man who you wanna be

		love the people who charming you
		second content here
	`}
	res, err := http.PostForm(apiUrl, form)
	if err != nil {
		fmt.Println("Http Post Failed with err:", err.Error(), " Check Server Log For Detail")
		t.Error(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read Body Error, err:", err.Error())
		t.Error(err.Error())
	}
	fmt.Println("Body:", string(body))
}

func TestPostHandler_DialysisConent(t *testing.T) {
	baseUrl := "http://localhost:3006/utils/extract/article?url="
	contentUrl := "https://www.huxiu.com/article/242750.html"
	res, err := http.Get(baseUrl + contentUrl)
	if err != nil {
		t.Error(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Got Response:\n%s", body)
}

func TestPostHandler_ArticleDetail(t *testing.T) {
	baseUrl := "http://localhost:3006/lens/article/detail?path=Hello-Kitty-Nice-To-meet-you.md&type=md"
	res, err := http.Get(baseUrl)
	if err != nil {
		t.Error(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Response:\n%s", string(body))
}

func TestPostHandler_ArticleDetail2(t *testing.T) {
	baseUrl := "http://localhost:3006/lens/article/detail?path=notfind.md&type=md"
	res, err := http.Get(baseUrl)
	if err != nil {
		t.Error(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Response:\n%s", string(body))
}

func TestPostHandler_UUIDNAME(t *testing.T) {
	uuidName, _ := uuid.NewV4()
	name := base64.RawURLEncoding.EncodeToString(uuidName[:])
	fmt.Println(name)
	fileName := fmt.Sprintf("%x", uuidName)
	fmt.Println(fileName)
}
