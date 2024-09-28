package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"os"
)

var gcm cipher.AEAD

func getGCM() (cipher.AEAD, error) {
	key := []byte(os.Getenv("AES_KEY"))

	// generate cipher block from key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// implement gcm
	gcm, err = cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm, nil
}

func AESEncrypt(plaintext string) (string, error) {
	text := []byte(plaintext)

	gcm, err := getGCM()
	if err != nil {
		return "", err
	}

	// allocate appropriate size for the nonce to seal and open
	nonce := make([]byte, gcm.NonceSize())

	// populate nonce with random values
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}

	// finally encrypt using gcm and nonce
	ciphertextBytes := gcm.Seal(nonce, nonce, text, nil)

	// encode bytes to base64
	ciphertext := base64.StdEncoding.EncodeToString(ciphertextBytes)

	return ciphertext, nil
}

func AESDecrypt(ciphertext string) (string, error) {
	// decode from base64 string to bytes
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	gcm, err := getGCM()
	if err != nil {
		return "", err
	}

	// the ciphertext is prepended by the nonce hence we separate it
	nonce := ciphertextBytes[:gcm.NonceSize()]
	ciphertextBytes = ciphertextBytes[gcm.NonceSize():]

	// finally decrypt using ciphertextBytes and nonce
	plaintextBytes, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	// convert bytes to string
	plaintext := string(plaintextBytes)

	return plaintext, nil
}
