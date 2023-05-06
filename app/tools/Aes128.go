package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var salt = "0123456789abcdef"

type Aes128 struct{}

func (c *Aes128) Encrypt(plainText string) (*string, error) {
	message := plainText

	key := []byte(salt) // 16-byte key for AES-128

	// initialize cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// initialize cipher mode
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	// encrypt message
	ciphertext := make([]byte, len(message))
	stream.XORKeyStream(ciphertext, []byte(message))

	// encode ciphertext to base64
	encrypted := base64.StdEncoding.EncodeToString(ciphertext)

	fmt.Println("Encrypted message:", encrypted)
	return &encrypted, nil
}

func (c *Aes128) Decrypt(encrypted string) (*string, error) {
	key := []byte(salt) // 16-byte key for AES-128

	// decode base64 encrypted message to ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	// initialize cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// initialize cipher mode
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	// decrypt ciphertext
	message := make([]byte, len(ciphertext))
	stream.XORKeyStream(message, ciphertext)

	result := string(message)

	fmt.Println("Decrypted message:", string(message))
	return &result, nil
}
