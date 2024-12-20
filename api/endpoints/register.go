package endpoints

import (
	"encoding/json"
	"jax/store"
	"net/http"

	"github.com/google/uuid"
)

type registerRequest struct {
	FirstName string `json:"first_name" validate:"required,min=10,max=200"`
	LastName  string `json:"last_name" validate:"required,min=10,max=200"`
	Username  string `json:"username" validate:"required,min=10,max=200"`
	Password  string `json:"password" validate:"required,min=10,max=200"`
}

type registerResponse struct {
	Token string `json:"token"`
}

func (e *Endpoint) Register(w http.ResponseWriter, r *http.Request) {
	req := registerRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// return invalid json
	}

	//validate(&req)

	pwd, err := e.hash.Make(req.Password)
	if err != nil {
		e.log.Error("failed to create the password hash", err)
		serverErrorResponse(w)
		return
	}

	u := store.NewUser{
		Uuid: uuid.NewString(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Password:  pwd,
	}

	err = e.db.AddUserAccount(&u)
	if err != nil {
		e.log.Error("failed to store user account", err)
		serverErrorResponse(w)
		return
	}

	// generate token
	token, err := e.createJwt(u.Uuid)
	if err != nil {
		e.log.Error("failed to create the jwt token", err)
		serverErrorResponse(w)
		return
	}

	registerSuccessResponse(w, token)
}

func registerSuccessResponse(w http.ResponseWriter, token string) {
	res := registerResponse{
		Token: token,
	}

	addHeaders(w)
	w.WriteHeader(http.StatusAccepted)

	_ = json.NewEncoder(w).Encode(res)
}

