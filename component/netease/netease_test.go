package netease

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestPaseNewsCategory(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	categorys, err := PaseNewsCategory()
	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		for k, v := range categorys {
			fmt.Printf("category:%s link:%s\n", k, v)
		}
	}
	if category, has := categorys["头条"]; has {
		if list, err := ParseCategoryContent("http:" + category); err == nil {
			for _, url := range list {
				fmt.Println(url)
			}
		}
	}
}
