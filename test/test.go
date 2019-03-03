package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type ThirdTokenData struct {
	AppId   uint32 `json:"app_id"`
	Timeout int64  `json:"timeout"`
	Nonce   int64  `json:"nonce"`
	IdName  string `json:"id_name"`
}

func GenerateThirdToken(appId uint32, idName string, serverSecret string, timeout int64) (string, error) {
	thirdTokenData := &ThirdTokenData{
		AppId:   appId,
		Timeout: timeout,
		Nonce:   1550646619,
		IdName:  idName,
	}

	plaintText, err := json.Marshal(thirdTokenData)
	if err != nil {
		return "", err
	}

	fmt.Println("plaintText:", string(plaintText), "s")
	fmt.Println("plaintText length:", len(string(plaintText)))

	iv := []byte("fe0b1ecb3d700a5f")
	fmt.Println("iv length:", len(iv))
	crypted, err := AesEncrypt(plaintText, []byte(serverSecret), iv)
	if err != nil {
		return "", err
	}

	// buf := string(iv) + string(crypted)
	fmt.Println("iv:", string(iv))
	buf := append(iv, crypted...)
	version := "01"
	thirdToken := version + base64.StdEncoding.EncodeToString([]byte(buf))
	return thirdToken, nil
}

func AesEncrypt(origData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	fmt.Println("blocksize:", blockSize)
	origData = AesPKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	fmt.Println("origData:", string(origData))
	fmt.Println("origData length:", len(origData))
	fmt.Println("crypted:", string(crypted))
	fmt.Println("crypted length:", len(crypted))
	return crypted, nil
}

func AesPKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func main() {
	token, _ := GenerateThirdToken(489705722, "360797139", "17550dd12092152ffe0b1ecb3d700a5f", 1550646619)
	fmt.Println("token:", token)
}
