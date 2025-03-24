package passhash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 12
)

// Generate makes password hash from password, using bcrypt algorithm.
// If an error is appeared it returns the empty string and the error.
func Generate(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// IsPasswordsMatch compares the hashed password and the password
// returning true if they are match and false otherwise. If an error is
// appeared it returns false and the error.
func IsPasswordsMatch(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
