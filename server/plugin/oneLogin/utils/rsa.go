package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
)

//# 生成 1024 位长度的私钥
//openssl genrsa -out private.key 1024
// openssl pkcs8 -topk8 -inform PEM -in private.key -outform pem -nocrypt -out pkcs8Private.key
//# 生成公钥
//openssl rsa -in pkcs8Private.key -pubout -out pkcs8Public.key

//RSA加密
func RSA_Encrypt(plainText []byte, publicRaw string) string {
	publicRaw = strings.Trim(publicRaw, "\n")
	if !strings.HasPrefix(publicRaw, "-----BEGIN PUBLIC KEY-----") {
		publicRaw = fmt.Sprintf("%s\n%s\n%s", "-----BEGIN PUBLIC KEY-----", publicRaw, "-----END PUBLIC KEY-----")
	}

	//pem解码
	block, _ := pem.Decode([]byte(publicRaw))
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}

	//返回密文
	//return cipherText
	return EncodeToString(cipherText)
}

//RSA解密
func RSA_Decrypt(cipherText []byte, privateRaw string) ([]byte, error) {
	privateRaw = strings.Trim(privateRaw, "\n")
	if !strings.HasPrefix(privateRaw, "-----BEGIN RSA PRIVATE KEY-----") {
		privateRaw = fmt.Sprintf("%s\n%s\n%s", "-----BEGIN RSA PRIVATE KEY-----", privateRaw, "-----END RSA PRIVATE KEY-----")
	}
	//pem解码
	block, _ := pem.Decode([]byte(privateRaw))
	//X509解码
	privateKey, err := genPriKey(block.Bytes, PKCS8)

	if err != nil {
		return nil, err
	}
	//对密文进行解密
	//plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return nil, err
	}
	//返回明文
	return plainText, nil
}

func Sha256WithRsa(msg string, privateRaw string) (string, error) {
	privateRaw = strings.Trim(privateRaw, "\n")
	if !strings.HasPrefix(privateRaw, "-----BEGIN RSA PRIVATE KEY-----") {
		privateRaw = fmt.Sprintf("%s\n%s\n%s", "-----BEGIN RSA PRIVATE KEY-----", privateRaw, "-----END RSA PRIVATE KEY-----")
	}

	blockPri, _ := pem.Decode([]byte(privateRaw))
	if blockPri == nil {
		return "", fmt.Errorf("blockPri is nil")
	}

	rsaPri, e := genPriKey(blockPri.Bytes, PKCS8)
	if e != nil {
		panic(e)
	}

	h := sha256.New()
	h.Write([]byte(msg))
	d := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPri, crypto.SHA256, d)
	if err != nil {
		return "", err
	}
	//encodedSig := base64.StdEncoding.EncodeToString(signature)
	//encodedSig := hex.EncodeToString(signature)

	encodedSig := EncodeToString(signature)
	return encodedSig, nil
}

const (
	PKCS1 int64 = iota
	PKCS8
)

func genPriKey(privateKey []byte, privateKeyType int64) (*rsa.PrivateKey, error) {
	var priKey *rsa.PrivateKey
	var err error
	switch privateKeyType {
	case PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
		}
	case PKCS8:
		{
			prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
			priKey = prkI.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, fmt.Errorf("unsupport private key type")
		}
	}
	return priKey, nil
}
