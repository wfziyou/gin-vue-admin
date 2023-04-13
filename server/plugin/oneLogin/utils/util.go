package utils

import (
	"bytes"
	"encoding/pem"
	"fmt"
	"os"
	"runtime"
)

// encodeURIComponent like same function in javascript
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent
//
// http://www.ecma-international.org/ecma-262/6.0/#sec-uri-syntax-and-semantics
func encodeURIComponent(s string, excluded ...[]byte) string {
	var b bytes.Buffer
	written := 0

	for i, n := 0, len(s); i < n; i++ {
		c := s[i]

		switch c {
		case '-', '_', '.', '!', '~', '*', '\'', '(', ')':
			continue
		default:
			// Unreserved according to RFC 3986 sec 2.3
			if 'a' <= c && c <= 'z' {

				continue

			}
			if 'A' <= c && c <= 'Z' {

				continue

			}
			if '0' <= c && c <= '9' {

				continue
			}
			if len(excluded) > 0 {
				conti := false
				for _, ch := range excluded[0] {
					if ch == c {
						conti = true
						break
					}
				}
				if conti {
					continue
				}
			}
		}

		b.WriteString(s[written:i])
		fmt.Fprintf(&b, "%%%02X", c)
		written = i + 1
	}

	if written == 0 {
		return s
	}
	b.WriteString(s[written:])
	return b.String()
}

// 填充最后一组
// plainText：明文
// blockSize：块大小
// 返回：填充后的明文
func PaddingLastGroup(plainText []byte, blockSize int) []byte {
	// get size num of last group, eg 28%8 = 4, 32%8=0
	padNum := blockSize - len(plainText)%blockSize
	// create a new slice, length equals padNum
	char := []byte{byte(padNum)}
	newPlain := bytes.Repeat(char, padNum)
	// append newPlain to plainText
	plainText = append(plainText, newPlain...)
	return plainText
}

// 去掉填充
// plainText：填充后的明文
// 返回：填充前的明文
func UnpaddingLastGroup(plainText []byte) []byte {
	length := len(plainText)
	// get length we need to cut
	number := int(plainText[length-1])
	return plainText[:length-number]
}

// 错误格式化
func Error(file string, line int, err string) error {
	return fmt.Errorf("file:%s line:%d error:%s", file, line, err)
}

// 读取公钥/私钥文件，获取解码的pem块
// filePath文件路径
// 返回pem块和错误
func GetKey(filePath string) (*pem.Block, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	fileInfo, err := file.Stat()
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, Error(file, line+1, err.Error())
	}
	block, _ := pem.Decode(buf)
	return block, err
}
