package uolo_center

import (
	"net/http"
	"net/url"
	"testing"
)

func TestUoloCenter_SaveProfile(t *testing.T) {
	form := url.Values{
		"id":       []string{"adfasdfasdfasdfasdfasfd"},
		"nickName": []string{"gonghuan"},
		"sex":      []string{"1"},
		"birthday": []string{"hello world"},
	}
	http.PostForm("https://api.echoface.cn/user/profile/save", form)
}
