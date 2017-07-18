package main

import (
	"crypto/sha1"
	"crypto/rsa"
	"crypto"
	"log"
	"encoding/base64"
	"sort"
	"fmt"
	"encoding/pem"
	"crypto/x509"
	"io"
	"encoding/hex"
)

/**
 * RSA签名
 * @param $data 待签名数据
 * @param $private_key_path 商户私钥文件路径
 * return 签名结果
 */
func RsaSign(origData string, privateKey *rsa.PrivateKey) (string, error) {

	h := sha1.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)

	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, digest)
	if err != nil {
		log.Println("rsaSign SignPKCS1v15 error", err)
		return "", err
	}
	data := base64.StdEncoding.EncodeToString(s)
	return string(data), nil
}

//签名字符串.
func AlipaySignString(mapBody map[string]interface{}) string {

	sorted_keys := make([]string, 0)
	for k, _ := range mapBody {
		sorted_keys = append(sorted_keys, k)
	}

	sort.Strings(sorted_keys)
	var signStrings string

	index := 0
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mapBody[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value
		}
		//最后一项后面不要&
		if index < len(sorted_keys)-1 {
			signStrings = signStrings + "&"
		}
		index++
	}

	return signStrings
}

func RSAVerify(src []byte, sign []byte) (pass bool, err error) {

	//步骤1，加载RSA的公钥
	block, _ := pem.Decode([]byte(GetAlipay().PublicKey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse RSA public key: %s\n", err)
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)

	//步骤2，计算代签名字串的SHA1哈希
	t := sha1.New()
	io.WriteString(t, string(src))
	digest := t.Sum(nil)

	//步骤3，base64 decode,必须步骤，支付宝对返回的签名做过base64 encode必须要反过来decode才能通过验证
	data, _ := base64.StdEncoding.DecodeString(string(sign))

	hexSig := hex.EncodeToString(data)
	fmt.Printf("base decoder: %v, %v\n", string(sign), hexSig)

	//步骤4，调用rsa包的VerifyPKCS1v15验证签名有效性
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
	if err != nil {
		fmt.Println("Verify sig error, reason: ", err)
		return false, err
	}

	return true, nil
}
