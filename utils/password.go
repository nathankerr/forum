package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hash from a plaintext string.
func HashPassword(plaintext string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// CheckHash compares the hash with a plaintext string.
func CheckHash(hash []byte, plaintext []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, plaintext)
	if err != nil {
		return err
	}
	return nil
}
