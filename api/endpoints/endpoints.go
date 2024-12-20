package endpoints

import (
	"encoding/json"
	"jax/hash"
	"jax/logger"
	"jax/store"
	"net/http"
)

type Endpoint struct {
	db store.Store
	hash hash.Hash
	log logger.Logger
}

func NewEndpoints(db store.Store, hash hash.Hash, log logger.Logger) *Endpoint {
	return &Endpoint{
		db: db,
		hash: hash,
		log: log,
	}
}

type registerRequest struct {
	FirstName string `json:"first_name" validate:"required,min=10,max=200"`
	LastName  string `json:"last_name" validate:"required,min=10,max=200"`
	Username  string `json:"username" validate:"required,min=10,max=200"`
	Password  string `json:"password" validate:"required,min=10,max=200"`
}

func (e *Endpoint) Register(w http.ResponseWriter, r *http.Request) {
	req := registerRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// return invalid json
	}

	//validate(&req)

	pwd, err := e.hash.Make(req.Password)
	if err != nil {
		panic("feck")
	}

	u := store.NewUser{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Username: req.Username,
		Password: pwd,
	}

	e.db.AddUserAccount(&u)
	// generate token
	// store toekn
	// return token


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

