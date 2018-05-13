package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestExtractArticle(t *testing.T) {
	mmap, _ := ExtractArticle("http://www.geekpark.net/news/229045")
	c, _ := json.Marshal(mmap)

	logrus.Println(string(c))
}
