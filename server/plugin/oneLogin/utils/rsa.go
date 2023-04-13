package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"strings"
)

//# 生成 1024 位长度的私钥
//openssl genrsa -out private.key 1024
//# 生成公钥
//openssl rsa -in private.key -pubout -out public-key.pem
// openssl pkcs8 -topk8 -inform PEM -in private.key -outform pem -nocrypt -out pkcs8Private.key

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
	encodedSig := hex.EncodeToString(signature)
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
