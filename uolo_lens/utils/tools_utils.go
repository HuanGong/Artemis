package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var ProvincePinYin map[string]string
var CityPinYin map[string]string

func GetCityByIp(ip string) string {
	const taobaoIp2GeoInfoUrl string = "http://ip.taobao.com/service/getIpInfo.php?ip="
	resp, err := http.Get(taobaoIp2GeoInfoUrl + ip)
	if err != nil {
		return "北京"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "北京"
	}
	type (
		Data struct {
			Country string `json:"country"`
			Region  string `json:"region"`
			City    string `json:"city"`
		}
		Res struct {
			Code int32 `json:"code"`
			Data Data  `json:"data"`
		}
	)
	r := &Res{}
	if err := json.Unmarshal(body, r); err != nil {
		return "北京"
	}
	if r.Code != 0 {
		return "北京"
	}

	return r.Data.City
}
