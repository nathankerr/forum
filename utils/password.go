package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plaintext string) ([]byte, error) {
	// Generate hash from plaintext.
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func CheckHash(hash []byte, plaintext []byte) error {
	// Compare hash with plaintext.
	err := bcrypt.CompareHashAndPassword(hash, plaintext)
	if err != nil {
		return err
	}
	return nil
}
