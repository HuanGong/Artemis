package utils

import (
	b64 "encoding/base64"
	"github.com/gogo/protobuf/proto"
)

func PBEncode(pb proto.Message) (string, error) {
	var str string
	value, err := proto.Marshal(pb)
	if err != nil {
		return "", err
	}

	str = b64.RawURLEncoding.EncodeToString(value)
	return str, err
}

func PBDecode(str string, pb proto.Message) error {

	data, _ := b64.RawURLEncoding.DecodeString(str)

	err := proto.Unmarshal(data, pb)
	if err != nil {
		return err
	}

	return nil
}
