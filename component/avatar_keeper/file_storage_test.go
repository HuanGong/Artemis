package avatar_keeper

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"testing"
)

var (
	storage = &FileStorage{
		Directory: ".",
	}
)

func TestFileStorage_GetFullPath(t *testing.T) {
	fmt.Println(storage.GetFullPath("gonghuan"))
}

func TestFileStorage_Store(t *testing.T) {
	p, _ := filepath.Abs(".")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugln("current dir ", p)

	data := &AvatarData{
		Small:  []byte("small"),
		Middle: []byte("middle"),
		Normal: []byte("nomal"),
	}
	storage.Store("gonghuan", data)
}

func TestFileStorage_Fetch(t *testing.T) {
	data, err := storage.Fetch("gonghuan", KAvatarSmall|KAvatarNormal|KAvatarMiddle)
	if err != nil {
		t.Error("fetch err ", err.Error())
	}
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debugln("small:", string(data.Small))
	logrus.Debugln("middle:", string(data.Middle))
	logrus.Debugln("normal:", string(data.Normal))
}

func TestFileStorage_Delete(t *testing.T) {
	storage.Delete("gonghuan")
}
