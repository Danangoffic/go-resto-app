package user

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$argon2id$v=%d,$m=%d,t=%d,p=%d$%s$%s"
)

func (r *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	argonHash := argon2.IDKey([]byte(password), salt, r.time, r.memory, r.threads, r.keyLen)

	b64Hash := r.encrypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, r.memory, r.time, r.threads, b64Salt, b64Hash)

	return encodedHash, nil
}

func (r *userRepo) encrypt(text []byte) string {
	nonce := make([]byte, r.gcm.NonceSize())

	cipherText := r.gcm.Seal(nonce, nonce, text, nil)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func (r *userRepo) decrypt(cipherText string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}

	if len(decoded) < r.gcm.NonceSize() {
		return nil, errors.New("Invalid NonceSize")
	}

	return r.gcm.Open(nil,
		decoded[:r.gcm.NonceSize()],
		decoded[r.gcm.NonceSize():],
		nil,
	)
}

func (r *userRepo) comparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")
	fmt.Printf("parts : %v\n", parts)
	var memory, time uint32
	var parallelism uint8

	switch parts[1] {
	case "argon2id":
		fmt.Printf("case : argon2id\n")
		fmt.Printf("parts3 : %v\n", parts[3])
		fmt.Printf("memory :%d, time:%d, parallelism:%d\n", &memory, &time, &parallelism)
		_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &parallelism)
		if err != nil {
			fmt.Printf("error with %v\n", err)
			return false, err
		}

		salt, err := base64.RawStdEncoding.DecodeString(parts[4])
		if err != nil {
			return false, err
		}

		hash := parts[5]

		decryptedHash, err := r.decrypt(hash)
		if err != nil {
			return false, err
		}

		var keyLen = uint32(len(decryptedHash))

		comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, parallelism, keyLen)

		return subtle.ConstantTimeCompare(comparisonHash, decryptedHash) == 1, nil
	default:
		fmt.Printf("unverified")
		return false, errors.New("Password didn't match\n")
	}

	// return false, nil
}
