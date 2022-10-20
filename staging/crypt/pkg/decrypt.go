package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
)

func Decrypt(key, filePath, fileType string) (string, error) {
	ciphertext, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Panic(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}
	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Panic(err)
	}
	tempPath := GetFilePath(filePath)
	deFileName := tempPath + "." + fileType
	if err := ioutil.WriteFile(deFileName, plaintext, 0777); err != nil {
		return "", err
	}
	return deFileName, nil
}
