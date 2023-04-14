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

func TestRsa() {
	PriKey := "MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBALSsLeIkMDoN0aWxoqK0sENGjNEC84nfHAFykEwWdJtXN+sJxAYkMo1A8e/hljt71YRL1QkYmZtek/Vi8XWhKD0Fme3rih50T3+Koe6YkLBKLEjM6MEmdV0z3lOmyCxbIYORuv5giiRUFSrAbKFFh0vUzezmyFefAAfqA316+eDhAgMBAAECgYAIt433LTvOcUA+OFXad9FRTaQZqYTKkCMvxrFDmonBvPGLu4rjqPdvbUS/CClRcWYZ3fbHW5J9tpB49G8l98KTKlTmbHWm2m6UBhqHxuNnhGK24muz99c1GzbXWgppxa2HeiA0yu2VT56RDfUrZGqvv+jD\nnpbuRP3YInXiMirF4QJBAOd7rH2ttqw6c1HHvL+9f3cWuUAXHcTqNhsWVlU2HslDCJX2VDW76lF+nZLp7xeNsHb8gxBOvcfd2YYs4tyiHa0CQQDHztosmZY/JDtJUtY+5xMMOWegCkeblCTLtc7v8zjLu2lcn6wEuShDXfxgVg7VPCTKbTiqYy+gEg9GsNKj\nxQ6FAkAMQ3EP93QGC9KwMnS9c7ydAoct7guVsxLKvJQ2T3eyEesShspPTnVLe/m9Hseb59XBd/85jfJf9FDh2t7p8WzBAkAKs4Nv3BH188TRGoSq/clBYFmycpp/NKH73xLkOwyRrMnp0gtufVQwt3nq1vEYbo4x4UOlrIZCdnUm/hVp/AXZAkAVhL9u6RGGM8TY4sDI5e9DZM/oyvM5R96uGF1Xb6LK7Znicb/P7eD1tWd0wCTI34e8iE7qxzfnNz02SMfMuWaE"
	publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC0rC3iJDA6DdGlsaKitLBDRozRAvOJ3xwBcpBMFnSbVzfrCcQGJDKNQPHv4ZY7e9WES9UJGJmbXpP1YvF1oSg9BZnt64oedE9/iqHumJCwSixIzOjBJnVdM95TpsgsWyGDkbr+YIokVBUqwGyhRYdL1M3s5shXnwAH6gN9evng4QIDAQAB"

	//PriKey := "MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAIC36ljhxVqx1DgpUC8ZRhpScOs9X3FESzSjcvLVv2SXpwuMEHJIj7RpByk7rJ4FVHfXl52cotMdDVWi7BRQQz9xhIxyJG/H4zeeoDQHjg7PHYFDTVdsN/DRHDv1tipOBny0IrrGWW677Sid6Z/sMeUpKPL/8+eWh7JwO1arUn3hAgMBAAECgYAOzJZ+F58oOU/sERvt/lroBdiDw2+oxzBaYfyCXP7/YsxK8JSnfx4+oOC45eqH1JcMnFYLQgoaebmhwfSgtUW15+sdBReIaciZgS/p5GNk+j7pxgP4N9AJQiXGHxG235pibBAEVH92qnBHgIpN40Hza1CWFhgTw1GOaFMPcz8qvQJBAOHGkbh4cn7bNxhslz0nvARjAMyekZdU4QCyQCCi0J3SfvxvzDRrvQVqun5AxUHTUNWhlRY3yub7DnYsFkqx3mcCQQCR8ybrZGMqNzeaxq/Fs5HTzbJosZ5k7nDfs5kSrW2NgY3+SJyJk+StrYvS4hZOXwg0iyiMrB1iwWmlaZkb8YR3AkBTauBwPeBfynLizUxbxhCLtmCXOYclWLEBZtqWtFFL3ngYoN3cCGqAU9yvxRKcrYzSQa8p1FddXCkNtGBQHMPFAkEAgCyNSn6QBBwYDipdZX+tGthzzTPnyfYJVLwyO0/pfTOA0wdLyhsC4nAd8qaxNkSJPTPU+a2R5Q+8yxLw7rRtQwJAeSj6/m5EOh+dCad0hyOZnzIO4Xaf8O9vujm6DeZ/smt3eLZis3OrlOwiReYb48R5N4OtUcJXMfghBFTf7QBMSw=="
	//publicKey := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCAt+pY4cVasdQ4KVAvGUYaUnDrPV9xREs0o3Ly1b9kl6cLjBBySI+0aQcpO6yeBVR315ednKLTHQ1VouwUUEM/cYSMciRvx+M3nqA0B44Ozx2BQ01XbDfw0Rw79bYqTgZ8tCK6xlluu+0onemf7DHlKSjy//PnloeycDtWq1J94QIDAQAB"

	str := "123456"
	sign, _ := Sha256WithRsa(str, PriKey)
	data := RSA_Encrypt([]byte(str), publicKey)
	aaa, _ := hex.DecodeString(data)
	newStr, _ := RSA_Decrypt(aaa, PriKey)
	fmt.Sprintf("%s,%s,%s", sign, data, newStr)
}
