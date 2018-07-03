package avatar_keeper

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/pkg/errors"
	"image/png"
)

type AvatarLevel int

const (
	KAvatarSmall  AvatarLevel = 0x1
	KAvatarMiddle AvatarLevel = 0x2
	KAvatarNormal AvatarLevel = 0x4
)

var (
	AvatarSmallSize  = [2]int{64, 64}
	AvatarMiddleSize = [2]int{150, 150}
	AvatarNormalSize = [2]int{400, 400}

	DefaultFileStoragePath = "./Avatar"
)

type (
	AvatarKeeper struct {
		storage Storage
	}

	AvatarData struct {
		Small  []byte
		Middle []byte
		Normal []byte
	}

	Storage interface {
		Delete(key string) error
		Store(key string, data *AvatarData) error
		Fetch(key string, level AvatarLevel) (*AvatarData, error)
	}
)

func NewAvatarKeeperWithFileStorage(baseDir string) *AvatarKeeper {
	s := &FileStorage{
		Directory: baseDir,
	}
	impl := &AvatarKeeper{
		storage: s,
	}
	return impl
}

func NewAvatarKeeper(s Storage) *AvatarKeeper {
	return &AvatarKeeper{
		storage: s,
	}
}

func (keeper *AvatarKeeper) SaveNewProfilePhoto(key string, buf *bytes.Buffer) error {
	if buf == nil {
		return errors.New("Invalid Image Buffer[nil]")
	}
	img, e := imaging.Decode(buf)
	if e != nil {
		return errors.Wrapf(e, "Decode Buffer to image Failed")
	}

	storageData := &AvatarData{}

	normal := imaging.Resize(img, AvatarNormalSize[0], AvatarNormalSize[1], imaging.Lanczos)
	normalBuf := new(bytes.Buffer)
	if err := png.Encode(normalBuf, normal); err != nil {
		return err
	}
	storageData.Normal = normalBuf.Bytes()

	middle := imaging.Resize(img, AvatarMiddleSize[0], AvatarMiddleSize[1], imaging.Lanczos)
	middleBuf := new(bytes.Buffer)
	if err := png.Encode(middleBuf, middle); err != nil {
		return err
	}
	storageData.Middle = middleBuf.Bytes()

	small := imaging.Resize(img, AvatarSmallSize[0], AvatarSmallSize[1], imaging.Lanczos)
	smallBuf := new(bytes.Buffer)
	if err := png.Encode(smallBuf, small); err != nil {
		return err
	}
	storageData.Small = smallBuf.Bytes()

	return keeper.storage.Store(key, storageData)
}

func (keeper *AvatarKeeper) GetSingleAvatar(key string, level AvatarLevel) ([]byte, error) {
	var singleLevel AvatarLevel
	switch level {
	case KAvatarSmall:
		singleLevel = KAvatarSmall
	case KAvatarNormal:
		singleLevel = KAvatarNormal
	default:
		singleLevel = KAvatarMiddle
	}
	data, err := keeper.storage.Fetch(key, singleLevel)
	if err != nil {
		return nil, err
	}

	switch singleLevel {
	case KAvatarSmall:
		return data.Small, nil
	case KAvatarMiddle:
		return data.Middle, nil
	case KAvatarNormal:
		return data.Normal, nil
	default:
		break
	}
	return nil, errors.New("should not reached")
}
