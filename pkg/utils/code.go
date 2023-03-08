package utils

import "encoding/base64"

func Pb2Base64(b []byte) (res string) {
	res = base64.StdEncoding.EncodeToString(b)
	return res
}

func Base642Pb(in string) (res []byte, err error) {
	res, err = base64.StdEncoding.DecodeString(in)
	return res, err
}
