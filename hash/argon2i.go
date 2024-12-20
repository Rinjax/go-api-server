package hash

import (
	"crypto/rand"
	"fmt"
	"jax/config"
	"golang.org/x/crypto/argon2"
)

type Argon2i struct {
	config *config.Argon2iConfig
}

func NewArgon2i(cfg *config.Argon2iConfig) *Argon2i {
	return &Argon2i{config: cfg}
}

func (a *Argon2i) Make(s string) (string, error) {
	// Generate a random salt
	salt := make([]byte, 16)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Hash the password using Argon2i
	hash := argon2.Key(
        []byte(s),
        salt,
        a.config.TimeCost,
        a.config.MemoryCost,
        a.config.Parallelism,
        a.config.HashLength,
    )

	// Return the salt + hash in a secure format
	return fmt.Sprintf("%x$%x", salt, hash), nil
}

func (a *Argon2i) Check(hash, str string) (bool, error) {
	// Split stored hash into salt and hash
	var salt, storedHashBytes []byte

	_, err := fmt.Sscanf(hash, "%x$%x", &salt, &storedHashBytes)
	if err != nil {
		return false, err
	}

	compareHash := argon2.Key(
        []byte(str),
        salt,
        a.config.TimeCost,
        a.config.MemoryCost,
        a.config.Parallelism,
        a.config.HashLength,
    )

	// Compare the newly computed hash with the stored hash
	return string(compareHash) == string(storedHashBytes), nil
}