package hash

import "jax/config"

type Hash interface {
	Make(s string) (string, error)
	Check(hash, str string) (bool, error)
}

func NewHash(cfg *config.Config) Hash {
	switch cfg.Hashing.Hasher {
	case "argon2i":
			return NewArgon2i(&cfg.Hashing.Argon2)
	default:
		panic("hasher not supported")
	}
}