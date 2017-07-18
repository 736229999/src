package zhongfu

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func desEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = pkcs5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func desDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

//------------------------------------------------------------------------------------------------------------------------
func encrypt(origData, key []byte) ([]byte, error) {
	ret, err := desEncrypt(origData, key)
	if err != nil {
		return nil, err
	}
	// 中福那边解密是UTF-8编码的字符串
	ret = []byte(fmt.Sprintf("%X", ret))
	return ret, nil
}

func decrypt(crypted, key []byte) ([]byte, error) {
	// 按刘伟给的代码加上这段……
	sData := string(crypted)
	buf := make([]byte, len(sData)/2)
	max := len(sData) / 2
	for x := 0; x < max; x++ {
		subStr := sData[x*2 : x*2+2]
		v, err := strconv.ParseInt(subStr, 16, 32)
		if err != nil {
			log.Panic(err)
		}
		buf[x] = byte(v)
	}
	return desDecrypt(buf, key)
}

//------------------------------------------------------------------------------------------------------------------------
func ZhongfuEncrypt(arg interface{}) ([]byte, error) {
	data, err := json.Marshal(arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return encrypt(data, []byte(KEY))
}

func ZhongfuDecrypt(data []byte, result interface{}) error {
	decrypted, err := decrypt(data, []byte(KEY))
	if err != nil {
		log.Println(err)
		return err
	}
	if err := json.Unmarshal(decrypted, result); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
