package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(message, key []byte) ([]byte, error) {
	// Create a new AES block cipher
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return nil, err
	}

	// Create a new GCM encryption mode
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create a random nonce
	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	// Encrypt the message using a nonce and propagate the tag
	ciphertext := aead.Seal(nil, nonce, message, nil)
	ciphertext = append(nonce, ciphertext...)
	return ciphertext, nil
}

func Decrypt(cipherText, key []byte) (string, error) {
	// Create a new AES block cipher
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return "", err
	}

	// Create a new GCM encryption mode
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract the nonce from the encrypted message
	nonceSize := aead.NonceSize()
	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]

	// Decode the message and check the tag
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", nil
	}

	return string(plaintext), nil
}
