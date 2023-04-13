package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	errorDes "github.com/flipped-aurora/gin-vue-admin/server/plugin/oneLogin/errors"
	"os"
	"runtime"
	"strings"
)

//# 生成 1024 位长度的私钥
//openssl genrsa -out private-key.pem 1024
//# 生成公钥
//openssl rsa -in private-key.pem -pubout -out public-key.pem

// 生成RSA密钥对
// keySize 密钥大小
// dirPath 密钥对文件路径
// 返回错误
func GenerateRsaKey(keySize int, dirPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	// x509
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// pem Block
	block := &pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}
	// just joint, caller must let dirPath right
	file, err := os.Create(dirPath + "private.pem")
	defer file.Close()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	err = pem.Encode(file, block)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	// get PublicKey from privateKey
	publicKey := privateKey.PublicKey
	derStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	block = &pem.Block{
		Type:  "rsa public key",
		Bytes: derStream,
	}
	file, err = os.Create(dirPath + "public.pem")
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	err = pem.Encode(file, block)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	return nil
}

// Rsa加密
// plainText 明文
// filePath 公钥文件路径
// 返回加密后的结果 错误
func RsaEncrypt(plainText []byte, filePath string) ([]byte, error) {
	// get pem.Block
	block, err := GetKey(filePath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	// X509
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	publicKey, flag := publicInterface.(*rsa.PublicKey)
	if flag == false {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, errorDes.RsatransError)
	}
	// encrypt
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	return cipherText, nil
}

// Rsa解密
// cipherText 密文
// filePath 私钥文件路径
// 返回解密后的结果 错误
func RsaDecrypt(cipherText []byte, filePath string) (plainText []byte, err error) {
	// get pem.Block
	block, err := GetKey(filePath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	// get privateKey
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	defer func() {
		if err2 := recover(); err2 != nil {
			_, file, line, _ := runtime.Caller(0)
			err = Error(file, line, errorDes.RsaNilError)
		}
	}()
	// get plainText use privateKey
	plainText, err3 := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err3 != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err3.Error())
	}
	return plainText, err
}

// Rsa签名
// plainText 明文
// filePath 私钥文件路径
// 返回签名后的数据 错误
func RsaSign(plainText []byte, priFilePath string) ([]byte, error) {
	// get pem.Block
	block, err := GetKey(priFilePath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	// calculate hash value
	hashText := sha512.Sum512(plainText)
	// Sign with hashText
	signText, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA512, hashText[:])
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	return signText, nil
}

// Rsa签名验证
// plainText 明文
// filePath 公钥文件路径
// 返回签名后的数据 错误
func RsaVerify(plainText []byte, pubFilePath string, signText []byte) error {
	// get pem.Block
	block, err := GetKey(pubFilePath)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	// x509
	pubInter, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	pubKey := pubInter.(*rsa.PublicKey)
	// hashText to verify
	hashText := sha512.Sum512(plainText)
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA512, hashText[:], signText)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return Error(file, line+1, err.Error())
	}
	return nil
}

const (
	PEM_BEGIN = "-----BEGIN RSA PRIVATE KEY-----\n"
	PEM_END   = "-----END RSA PRIVATE KEY-----"
)

//RsaSign Rsa签名
//signContent: 签名内容
//privateKey：私钥字符串，PKCS#1，开头去除-----BEGIN RSA PRIVATE KEY-----，结尾去掉-----END RSA PRIVATE KEY-----
//hash: hash算法
func RsaSignEx(signContent string, privateKey string, hash crypto.Hash) ([]byte, error) {
	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)
	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func ParsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateKey = FormatPrivateKey(privateKey)
	// 2、解码私钥字节，生成加密对象
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	// 3、解析DER编码的私钥，生成私钥对象
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}
func FormatPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, PEM_BEGIN) {
		privateKey = PEM_BEGIN + privateKey
	}
	if !strings.HasSuffix(privateKey, PEM_END) {
		privateKey = privateKey + PEM_END
	}
	return privateKey
}

func SHA256withRSA(signContent string, privateKey string) (sign string, err error) {
	signData, err := RsaSignEx(signContent, privateKey, crypto.SHA256)
	if err != nil {
		return "", err
	}
	//return base64.StdEncoding.EncodeToString(signData)
	return hex.EncodeToString(signData), err
}
