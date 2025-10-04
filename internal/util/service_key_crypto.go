package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

const (
	androidUUID = "4025d4df-6e1d-4b2b-a3dd-c6e009780010"
	iosUUID     = "6aff4b82-0454-4e41-982d-4251791c712d"
	webUUID     = "4f91e092-ff67-4fff-8516-6f51ee9e2a4f"

	androidServiceKey = "ElJJwfCnMkcJvUnzspexarKL2auKLycZEM6xoDr8yiQJkqzwHbOCwUk3GyKpvA.KkjYP0xozjOrkbJwCQ7QVg"
	iosServiceKey     = "E-ldOBKcGCSJ0W-KCTT6da6haJselZwnf4VQOGS-mpct5-kLwlsReKeaN8MADA.hZLa3MmavaVO7FJkHCdeoQ"
	webServiceKey     = "6uL46e_8SLokrBBQpxrSkAuuji2uFJQ7NIqwpDKWkj9TbNql2wxz-1Zxm-qKLw.o0SBcISsLnXSJetpE5ACqw"
)

var ENCRYPTION_KEY = []byte("api-secret-token-random-string--")

func encrypt(payload string) (string, error) {
	// IV (초기화 벡터) 생성
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %v", err)
	}

	block, err := aes.NewCipher(ENCRYPTION_KEY)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// CTR 암호화 스트림 생성
	stream := cipher.NewCTR(block, iv)
	plaintext := []byte(payload)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// 결과를 base64url로 인코딩
	encrypted := base64.RawURLEncoding.EncodeToString(ciphertext)
	ivEncoded := base64.RawURLEncoding.EncodeToString(iv)
	return fmt.Sprintf("%s.%s", encrypted, ivEncoded), nil
}

// Decrypt는 암호화된 payload를 복호화합니다.
func decrypt(encryptedPayload string) (string, error) {
	// 암호화된 텍스트와 IV를 "." 기준으로 분리
	parts := bytes.Split([]byte(encryptedPayload), []byte("."))
	if len(parts) != 2 {
		return "", errors.New("invalid payload format")
	}

	encryptedText, ivText := parts[0], parts[1]

	// Base64url 디코딩
	ciphertext, err := base64.RawURLEncoding.DecodeString(string(encryptedText))
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %v", err)
	}

	iv, err := base64.RawURLEncoding.DecodeString(string(ivText))
	if err != nil {
		return "", fmt.Errorf("failed to decode IV: %v", err)
	}

	block, err := aes.NewCipher(ENCRYPTION_KEY)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// CTR 복호화 스트림 생성
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}
