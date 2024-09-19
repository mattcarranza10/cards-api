package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

var (
	service *AES256EncryptionService
)

type AES256EncryptionService struct {
	secretKey []byte
}

func newAES256EncryptionService(secretKey []byte) *AES256EncryptionService {
	return &AES256EncryptionService{
		secretKey: secretKey,
	}
}

func newService() {
	secretKey, err := generateAES256SecretKey()
	if err != nil {
		log.Fatalf("Error generating encryption service secret key: %v", err)
	}
	service = newAES256EncryptionService(secretKey)
}

func generateAES256SecretKey() ([]byte, error) {
	return generateSecretKey(32)
}

func generateSecretKey(size int) ([]byte, error) {
	key := make([]byte, size)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func GetAES256EncryptionService() *AES256EncryptionService {
	if service == nil {
		newService()
	}
	return service
}

func (svc *AES256EncryptionService) Encrypt(data string) (string, error) {
	block, err := aes.NewCipher(svc.secretKey)
	if err != nil {
		return "", err
	}

	dst := make([]byte, aes.BlockSize+len(data))
	iv := dst[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(dst[aes.BlockSize:], []byte(data))

	return hex.EncodeToString(dst), nil
}

func (svc *AES256EncryptionService) Decrypt(encryptedData string) (string, error) {
	src, err := hex.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(svc.secretKey)
	if err != nil {
		return "", err
	}

	if len(src) < aes.BlockSize {
		return "", err
	}

	iv := src[:aes.BlockSize]
	src = src[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(src, src)

	return string(src), nil
}
