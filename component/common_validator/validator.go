package common_validator

import (
	"regexp"
)

// 域名合法性测试
func ValidateDomain(str string) bool {
	reg := `^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-\.\{\},@?^=%&:\/~\+#]*[\w\-\@?^=%&\/~\+#])?$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}

//校验Ip正确性
func ValidateIp(ip string) bool {
	reg := `^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[1-9])\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(ip)
}

//校验设备号
func ValidateDevices(device string) bool {
	reg := `^[0-9a-zA-Z_-]{1,}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(device)
}