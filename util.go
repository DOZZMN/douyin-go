package douyinGo

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/DOZZMN/douyin-go/auth"
)

// ---------------------------------------------------------------------------------------

type transport struct {
	http.RoundTripper
	credentials *auth.Credentials
}

func newTransport(credentials *auth.Credentials, tr http.RoundTripper) *transport {
	if tr == nil {
		tr = http.DefaultTransport
	}
	return &transport{tr, credentials}
}

func Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func Base64Decode(encodeString string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodeString)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != block.BlockSize() {
		var keySizeError aes.KeySizeError
		return nil, errors.New(keySizeError.Error())
	}
	encrypter := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	encrypter.CryptBlocks(origData, crypted)
	return PKCS5UnPadding(origData), nil
}
