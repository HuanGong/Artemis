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
	//apiUrl := "https://api.echoface.cn/lens/lens/article/new"

	form := url.Values{}

	form["title"] = []string{"Welcome Uolo World"}
	form["tag"] = []string{"C++"}
	form["author"] = []string{"Huan.Gong"}
	form["mime"] = []string{"md"}
	form["origin"] = []string{"http://www.baidu.com/news/abc"}
	form["summary"] = []string{"UOLO, You Only Live Once and Seven times have I despised my soul"}
	form["content"] = []string{`# UOLO, You Only Live Once

- do the things you wanna do
    + 做爱做的事情

- to be the man who you wanna be
    + 做想做的人

### Seven times have I despised my soul:

From: Kahlil Gibran

- The first time when I saw her being meek that she might attain height.
- The second time when I saw her limping before the crippled.
- The third time when she was given to choose between the hard and the easy, and she chose the easy.
- The fourth time when she committed a wrong, and comforted herself that others also commit wrong.
- The fifth time when she forbore for weakness, and attributed her patience to strength.
- The sixth time when she despised the ugliness of a face, and knew not that it was one of her own masks.
- And the seventh time when she sang a song of praise, and deemed it a virtue`}
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
	//contentUrl := "https://www.huxiu.com/article/242750.html"
	contentUrl := "http://start.iresearch.cn/content/2018/03/273475.shtml"
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
	baseUrl := "http://localhost:3006/lens/article/detail?path=20180506/J7u7SNeqQ5GPmZLiUOtfMA.md&type=md"
	res, err := http.Get(baseUrl)
	if err != nil {
		t.Error(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println("Response:", string(body))
}

func TestPostHandler_ArticleDetail2(t *testing.T) {
	//baseUrl := "http://localhost:3006/lens/article/detail?path=notfind.md&type=md"
	baseUrl := "http://api.echoface.cn/lens/lens/article/detail?path=notfind.md&type=md"

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
