package endpoints

import (
	"encoding/json"
	"net/http"
)

// headers map of common headers to be set on all responses
var headers = map[string]string{
	"Content-Type": "application/json",
}

type apiError struct {
	Err        any `json:"err"`
	StatusCode int `json:"status_code"`
}

// addHeaders function which writes the common headers map to the http response
func addHeaders(w http.ResponseWriter) {
	for h, v := range headers {
		w.Header().Set(h, v)
	}
}

func serverErrorResponse(w http.ResponseWriter) {
	res := apiError{
		Err: "oh bugger, somethings gone wrong!!",
		StatusCode: http.StatusInternalServerError,
	}

	addHeaders(w)
	w.WriteHeader(res.StatusCode)

	_ = json.NewEncoder(w).Encode(res)
}