package utils

import (
	"bytes"
	"encoding/base64"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	CleanMarkPath = "clean-mark"
)

func SaveArticleAsFileSync(fullPath, content string) error {
	folder, _ := filepath.Split(fullPath)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if fail := os.MkdirAll(folder, 0755); fail != nil {
			return fail
		}
	}
	return ioutil.WriteFile(fullPath, []byte(content), 0644)
}

func ExtractArticle(url string) (map[string]string, error) {

	result := make(map[string]string)
	if len(url) == 0 {
		return result, errors.New("Url Not Correct")
	}

	cmd := exec.Command(CleanMarkPath, url)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return result, errors.Wrapf(err, "cmd execute error")
	}

	const SHeaderStart = 0
	const SExtractHeader = 1
	const SExtractContent = 2

	state := SHeaderStart

EXTRACTLOOP:
	for {
		switch state {
		case SHeaderStart:
			line, err := out.ReadString('\n')
			if err != nil || io.EOF == err {
				return result, errors.Wrapf(err, "cmd not get content")
				break EXTRACTLOOP
			}
			if line == "---\n" {
				state = SExtractHeader
			}
			continue EXTRACTLOOP

		case SExtractHeader:

			line, err := out.ReadString('\n')

			if err != nil || io.EOF == err {
				return result, errors.Wrapf(err, "cmd not get content")
				break EXTRACTLOOP
			}
			if line == "---\n" {
				state = SExtractContent
				continue EXTRACTLOOP
			}

			heasers := strings.Split(line, ":")
			if len(heasers) < 2 {
				continue EXTRACTLOOP
			}
			switch heasers[0] {
			case "link":
				result["link"] = strings.Join(heasers[1:], ":")
			case "title":
				result["title"] = strings.Join(heasers[1:], ":")
			case "date":
				result["date"] = strings.Join(heasers[1:], ":")
			case "description":
				result["desc"] = strings.Join(heasers[1:], ":")
			case "author":
				result["author"] = strings.Join(heasers[1:], ":")
			case "keywords":
				result["keywords"] = strings.Join(heasers[1:], ":")
			case "publisher":
				result["publisher"] = strings.Join(heasers[1:], ":")
			default:
				continue EXTRACTLOOP
			}

		case SExtractContent:
			content := out.String()
			result["content"] = content
			break EXTRACTLOOP
		default:
			break EXTRACTLOOP
		}
	}
	return result, nil
}

func GenUoloUUID() string {
	uuidName, _ := uuid.NewV4()
	articleUID := base64.RawURLEncoding.EncodeToString(uuidName[:])
	return articleUID
}

func GenUUID() string {
	uuidGen, _ := uuid.NewV4()
	return uuidGen.String()
}
