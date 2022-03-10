package edcry

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5(raw string, salt string) string {
	h := md5.New()
	h.Write([]byte(raw + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(raw string, salt string) string {
	h := sha256.Sum256([]byte(raw + salt))
	return hex.EncodeToString(h[:])
}

// GenAesKey 生成 AES 密钥, 必须是 16 的倍数
func GenAesKey(width int) string {
	bytes := make([]byte, width)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	return hex.EncodeToString(bytes)
}
func GenAesKey16() string {
	return GenAesKey(16)
}
func GenAesKey32() string {
	return GenAesKey(32)
}

// AesEncrypt use counter mode
// ref: https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
func AesEncrypt(key, plaintext string) string {
	keybs, _ := hex.DecodeString(key)
	plaintextbs := []byte(plaintext)

	block, err := aes.NewCipher(keybs)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintextbs, nil)
	return fmt.Sprintf("%x", ciphertext)
}
func AesDecrypt(key, ciphertext string) string {
	keybs, _ := hex.DecodeString(key)
	enc, _ := hex.DecodeString(ciphertext)

	block, err := aes.NewCipher(keybs)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertextbs := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertextbs, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
