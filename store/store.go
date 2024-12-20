package store

import "jax/config"

type Store interface {
	AddUserAccount(u *NewUser) error
}

func NewStore(cfg *config.Config) Store {
	switch cfg.Store.Storer {
	case "faker":
		return NewFaker()
	default:
		panic("store not supported")
	}
}