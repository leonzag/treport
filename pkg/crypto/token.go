package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

func EncryptToken(pwd string, token string) (string, error) {
	encryptedTokenBytes, err := encrypt([]byte(pwd), []byte(token))
	if err != nil {
		return "", err
	}
	encryptedToken := hex.EncodeToString(encryptedTokenBytes)

	return encryptedToken, nil
}

func DecryptToken(pwd string, encryptedToken string) (string, error) {
	encryptedTokenBytes, err := hex.DecodeString(encryptedToken)
	if err != nil {
		return "", err
	}
	tokenBytes, err := Decrypt([]byte(pwd), encryptedTokenBytes)
	if err != nil {
		return "", err
	}

	return string(tokenBytes), nil
}

func encrypt(key, data []byte) ([]byte, error) {
	key, salt, err := deriveKey(key, nil)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func Decrypt(key, data []byte) ([]byte, error) {
	data, salt := data[:len(data)-32], data[len(data)-32:]

	key, _, err := deriveKey(key, salt)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func deriveKey(password, salt []byte) ([]byte, []byte, error) {
	keyLen, saltLen := 32, 32

	if salt == nil {
		salt = make([]byte, saltLen)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	nCost := 32768
	rCost, pCost := 8, 1
	key, err := scrypt.Key(password, salt, nCost, rCost, pCost, keyLen)
	if err != nil {
		return nil, nil, err
	}

	return key, salt, nil
}
