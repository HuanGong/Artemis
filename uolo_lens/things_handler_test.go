package uolo_lens

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestThingsHandler_NewThings(t *testing.T) {
	apiUrl := "http://localhost:3006/things/v1/new"
	//apiUrl := "https://api.echoface.cn/lens/article/new"

	form := url.Values{}

	form["content"] = []string{"Welcome to UoloThings"}
	form["comment"] = []string{"我希望能很好的完成它！"}

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

func TestThingsHandler_GetDoneThingsList(t *testing.T) {
	apiUrl := "http://localhost:3006/things/v1/done/list"

	res, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Http Post Failed with err:", err.Error(), " Check Server Log For Detail")
		t.Error(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("Body:", string(body))
}

func TestThingsHandler_GetThingsList(t *testing.T) {
	apiUrl := "http://localhost:3006/things/v1/list"

	res, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Http Post Failed with err:", err.Error(), " Check Server Log For Detail")
		t.Error(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("Body:", string(body))
}

func TestThingsHandler_MarkUoloThingFinish(t *testing.T) {
	apiUrl := "http://localhost:3006/things/v1/finish"
	form := url.Values{}

	form["uuid"] = []string{"785b7e88-ce3d-432f-9b22-11290e488cc3"}
	form["comment"] = []string{"I Finish A Uolo Things!"}
	res, err := http.PostForm(apiUrl, form)
	if err != nil {
		fmt.Println("Http Post Failed with err:", err.Error(), " Check Server Log For Detail")
		t.Error(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read Response Body Failed With:", err.Error())
		return
	}

	fmt.Println("Body:", string(body))
}

func TestThingsHandler_DeleteUoloThing(t *testing.T) {
	apiUrl := "http://localhost:3006/things/v1/delete"
	form := url.Values{}

	form["uuid"] = []string{"785b7e88-ce3d-432f-9b22-11290e488cc3"}
	res, err := http.PostForm(apiUrl, form)
	if err != nil {
		fmt.Println("Http Post Failed with err:", err.Error(), " Check Server Log For Detail")
		t.Error(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read Response Body Failed With:", err.Error())
		return
	}

	fmt.Println("Body:", string(body))
}
