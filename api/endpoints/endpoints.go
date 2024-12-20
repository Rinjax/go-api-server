package endpoints

import (
	"jax/config"
	"jax/hash"
	"jax/logger"
	"jax/store"
	"net/http"
)

type Endpoint struct {
	config *config.Config
	db store.Store
	hash hash.Hash
	log logger.Logger
}

func NewEndpoints(cfg *config.Config, db store.Store, hash hash.Hash, log logger.Logger) *Endpoint {
	return &Endpoint{
		config: cfg,
		db: db,
		hash: hash,
		log: log,
	}
}



func (e *Endpoint) Login(w http.ResponseWriter, r *http.Request) {
	// validate request
	// check if request has already has token
	// check user/pass db
	// generate token
	// store
	// return
}

func (e *Endpoint) Logout(w http.ResponseWriter, r *http.Request) {
	// validate request
	// check if request has already has token
	// check user/pass db
	// generate token
	// store
	// return
}

