package utils

import (
	"github.com/coreos/etcd/pkg/ioutil"
	"github.com/sirupsen/logrus"
)

func SaveArticleAsFileSync(fullPath, content string) {
	err := ioutil.WriteAndSyncFile(fullPath, []byte(content), 0644)
	if err != nil {
		logrus.Errorf("Save Content To file %s Failed, Content:%s", fullPath, content)
		//TODO: note down failed  content info to file for avoiding data loss
	}
}
