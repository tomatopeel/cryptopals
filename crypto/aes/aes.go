package aes

import (
	"crypto/aes"
)

func DecryptCbc(ciphertext, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	kl := len(key)
	tl := len(ciphertext)
	mod := tl % kl
	if mod != 0 {
		padding := Pkcs7Padding(ciphertext, kl-mod)
		for _, block := range padding {
			ciphertext = append(ciphertext, block)
		}
		tl = len(ciphertext)
	}
	decrypted := make([]byte, tl)

	for a, b := 0, kl; a < tl; a, b = a+kl, b+kl {
		cipher.Decrypt(decrypted[a:b], ciphertext[a:b])
		for i, j := a, 0; j < len(iv); i, j = i+1, j+1 {
			decrypted[i] ^= iv[j]
		}
		iv = ciphertext[a:b]
	}
	return decrypted, nil
}

func EncryptCbc(plaintext, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	kl := len(key)
	tl := len(plaintext)
	mod := tl % kl
	if mod != 0 {
		padding := Pkcs7Padding(plaintext, kl-mod)
		for _, block := range padding {
			plaintext = append(plaintext, block)
		}
		tl = len(plaintext)
	}
	encrypted := make([]byte, tl)

	for a, b := 0, kl; a < tl; a, b = a+kl, b+kl {
		for i, j := a, 0; j < len(iv); i, j = i+1, j+1 {
			plaintext[i] ^= iv[j]
		}
		cipher.Encrypt(encrypted[a:b], plaintext[a:b])
		iv = encrypted[a:b]
	}
	return encrypted, nil
}

func DecryptEcb(ciphertext, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	kl := len(key)        // key length
	tl := len(ciphertext) // total length
	decrypted := make([]byte, tl)

	// a = block start; b = block end
	for a, b := 0, kl; a < tl; a, b = a+kl, b+kl {
		cipher.Decrypt(decrypted[a:b], ciphertext[a:b])
	}

	return decrypted, nil
}

func EncryptEcb(plaintext, key []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	kl := len(key)
	tl := len(plaintext)
	encrypted := make([]byte, len(plaintext))

	for a, b := 0, kl; a < tl; a, b = a+kl, b+kl {
		cipher.Encrypt(encrypted[a:b], plaintext[a:b])
	}

	return encrypted, nil
}

func Pkcs7Padding(data []byte, padding int) []byte {
	blocks := make([]byte, padding)
	for i := 0; i < len(blocks); i++ {
		blocks[i] = byte(4)
	}
	return blocks
}
