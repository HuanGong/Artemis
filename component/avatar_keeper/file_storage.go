package avatar_keeper

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spaolacci/murmur3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type (
	FileStorage struct {
		Directory string
	}
	NameSet struct {
		small  string
		middle string
		normal string
	}
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (storage *FileStorage) GetFullPath(key string) (string, *NameSet) {

	l, h := murmur3.Sum128([]byte(key))

	lResult := make([]byte, 8)
	binary.LittleEndian.PutUint64(lResult, uint64(l))

	hResult := make([]byte, 8)
	binary.LittleEndian.PutUint64(hResult, uint64(h))

	first := hex.EncodeToString(lResult[:4])
	second := hex.EncodeToString(lResult[4:8])
	third := hex.EncodeToString(hResult[:4])
	last := hex.EncodeToString(hResult[4:8])

	name := murmur3.Sum32([]byte(key))
	names := &NameSet{
		small:  fmt.Sprintf("%d_small.png", name),
		middle: fmt.Sprintf("%d_middle.png", name),
		normal: fmt.Sprintf("%d_normal.png", name),
	}

	return filepath.Join(storage.Directory, first, second, third, last), names
}

func (storage *FileStorage) Store(key string, data *AvatarData) error {
	path, names := storage.GetFullPath(key)

	logrus.Debugln("storage data to path, abs", path)

	if f, err := os.Stat(path); err != nil {
		if err := os.MkdirAll(path, 0755); err != nil {
			return errors.Wrapf(err, "存储失败")
		}
	} else if !f.IsDir() {
		return errors.New("FileStorage Error")
	}

	if len(data.Small) > 0 {
		ioutil.WriteFile(filepath.Join(path, names.small), data.Small, 0644)
	}
	if len(data.Middle) > 0 {
		ioutil.WriteFile(filepath.Join(path, names.middle), data.Middle, 0644)
	}
	if len(data.Normal) > 0 {
		ioutil.WriteFile(filepath.Join(path, names.normal), data.Normal, 0644)
	}
	return nil
}

func (storage *FileStorage) Delete(key string) error {
	path, names := storage.GetFullPath(key)

	if hasFolder, _ := PathExists(path); !hasFolder {
		return nil
	}

	smallPath := filepath.Join(path, names.small)
	if exist, _ := PathExists(smallPath); exist {
		if e := os.Remove(smallPath); e != nil {
			logrus.Debugf("Remove File %s from path %s failed with %s", path, names.small, e.Error())
			return e
		}
	}
	middlePath := filepath.Join(path, names.middle)
	if exist, _ := PathExists(middlePath); exist {
		if e := os.Remove(middlePath); e != nil {
			logrus.Debugf("Remove File %s from path %s failed with %s", path, names.middle, e.Error())
			return e
		}
	}

	normalPath := filepath.Join(path, names.normal)
	if exist, _ := PathExists(normalPath); exist {
		if e := os.Remove(normalPath); e != nil {
			logrus.Debugf("Remove File %s from path %s failed with %s", path, names.middle, e.Error())
			return e
		}
	}

	return nil
}

func (storage *FileStorage) Fetch(key string, level AvatarLevel) (*AvatarData, error) {
	path, names := storage.GetFullPath(key)

	data := &AvatarData{}

	if level&KAvatarSmall != 0 {
		var e1 error
		if data.Small, e1 = ioutil.ReadFile(filepath.Join(path, names.small)); e1 != nil {
			logrus.Debugf("Fetch File %s from path %s failed with %s", path, names.small, e1.Error())
			return nil, errors.Wrapf(e1, "read small avatar failed")
		}
	}
	if level&KAvatarMiddle != 0 {
		var e1 error
		if data.Middle, e1 = ioutil.ReadFile(filepath.Join(path, names.middle)); e1 != nil {
			logrus.Debugf("Fetch File %s from path %s failed with %s", path, names.normal, e1.Error())
			return nil, errors.Wrapf(e1, "read middle avatar failed")
		}
	}
	if level&KAvatarNormal != 0 {
		fmt.Println("read normal ", level&KAvatarNormal != 0)
		var e1 error
		if data.Normal, e1 = ioutil.ReadFile(filepath.Join(path, names.normal)); e1 != nil {
			logrus.Debugf("Fetch File %s from path %s failed with %s", path, names.normal, e1.Error())
			return nil, errors.Wrapf(e1, "read normal avatar failed")
		}
	}
	return data, nil
}
