package utils

import (
	"bytes"
	"github.com/coreos/etcd/pkg/ioutil"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os/exec"
	"strings"
)

func SaveArticleAsFileSync(fullPath, content string) {
	err := ioutil.WriteAndSyncFile(fullPath, []byte(content), 0644)
	if err != nil {
		logrus.Errorf("Save Content To file %s Failed, Content:%s", fullPath, content)
		//TODO: note down failed  content info to file for avoiding data loss
	}
}

func ExtractArticleFromUrl(url string) (map[string]string, error) {

	result := make(map[string]string)

	cmd := exec.Command("clean-mark", url)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return result, errors.Wrapf(err, "cmd execute error")
	}

	result["content"] = out.String()

	metaStart := false
	for {
		line, err := out.ReadString('\n')
		if err != nil || io.EOF == err {
			return result, errors.Wrapf(err, "cmd not get content")
			break
		}

		//if strings.HasPrefix(line, "---") {
		if line == "---\n" {
			metaStart = !metaStart
			if metaStart == false {
				break
			}
			continue
		}

		if false == metaStart {
			continue
		}
		Infos := strings.Split(line, ":")
		if len(Infos) < 2 {
			continue
		}
		switch Infos[0] {
		case "link":
			result["link"] = strings.Join(Infos[1:], ":")
		case "title":
			result["title"] = strings.Join(Infos[1:], ":")
		case "description":
			result["desc"] = strings.Join(Infos[1:], ":")
		case "keywords":
			result["keywords"] = strings.Join(Infos[1:], ":")
		case "author":
			result["author"] = strings.Join(Infos[1:], ":")
		case "date":
			result["date"] = strings.Join(Infos[1:], ":")
		case "publisher":
			result["publisher"] = strings.Join(Infos[1:], ":")
		default:
			continue
		}
	}
	return result, nil
}
