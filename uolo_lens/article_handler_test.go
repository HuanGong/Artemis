package uolo_lens

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/JesusIslam/tldr"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestPostHandler_ArticleNew(t *testing.T) {
	//apiUrl := "http://localhost:3006/lens/article/new"
	apiUrl := "https://api.echoface.cn/lens/article/new"

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
	baseUrl := "http://localhost:3006/article/detail?path=20180506/J7u7SNeqQ5GPmZLiUOtfMA.md&type=md"
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
	baseUrl := "http://localhost:3006/article/detail?path=notfind.md&type=md"
	//baseUrl := "http://api.echoface.cn/lens/lens/article/detail?path=notfind.md&type=md"

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

func TestPostHandler_ArticlePartJson(t *testing.T) {
	type Response struct {
		Abc     string `json:"-"`
		Code    int32  `json:"code"`
		Message string
	}
	r := &Response{
		Abc:     "Hello",
		Code:    1,
		Message: "HuanGong",
	}
	s, _ := json.Marshal(r)
	logrus.Infoln("jsonstring:", string(s))
}

func TestAutoSummary(t *testing.T) {
	intoSentences := 3
	text := `继前万达人力总监张洪举加盟维基链之后，近日阿里巴巴商业分析师高航以数百万年薪加入维基链团队，担任高级副总裁和首席数据分析师，负责对维基链商业模式设计及战略决策提供技术支持，对产品和运营搭建数字化监控体系，并利用自身国际化背景帮助维基链完成海外推广工作。高航先生将于5月12日正式入职。
	高智商人才门萨俱乐部成员助力维基链

	高航先生毕业于常春藤名校哥伦比亚大学，曾是华尔街Fishbowl公司首席数据科学家，能力出众，是门萨国际俱乐部高级成员，阿里巴巴数据分析师。跟张洪举先生一样，高航对于区块链的运营模式非常感兴趣，在维基链诞生后高航先生高度称赞了维基链的商业模式，在拒绝了腾讯和华为之后主动从阿里巴巴离职加入维基链。高航先生将用其敏锐的商业嗅觉和出色的商业分析能力帮助维基链从区块链行业大数据分析、维基链客户群大数据等多方面提升效率，归纳总结深层次规律，为公司经营决策提供战略部署支援。在高航加盟之后维基链整个商业模式都会得到优化。

	维基链为何如此吸引人才

	维基链是利用区块链与传统商业结合，利用区块链的去中心化特性解决传统模式的痛点，目前主打游戏竞猜行业，利用区块链的特性天然解决了传统竞猜行业存在的问题，前景远大。维基链团队重视人才、渴求人才，曾多次表示，区块链未来的竞争也必然是人才的竞争。每个加入维基链核心团队的人都会得到团队的重视，年薪往往也是七位数以上，维基链一直对有能力的人才敞开大门，这一直吸引了多领域的人才加入。

	区块链项目的“人才争夺战”日益白热化，维基链在业内已实现阶段性领先。近期维基链主网上线，6月份世界杯开始，之后的维基链前景不可限量。

	5月13日13:00，维基链将于深圳举办竞猜产品发布会全球首站。据维基链官方消息，此次发布会的规模可达到千人以上。`

	bag := tldr.New()
	result, _ := bag.Summarize(text, intoSentences)
	fmt.Println("summary:", result)
}
