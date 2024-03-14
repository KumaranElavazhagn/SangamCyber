package handler

import (
	"encoding/json"
	"net/http"
	Service "sangamCyber/Service"
	"sangamCyber/dto"
	"sangamCyber/errs"
	"sangamCyber/validator"
	"strings"
)

type Handlers struct {
	Service Service.Service
}

// The `InsertUserInfo` function in the code snippet is a handler function that processes a request to
// insert user information. Here's a breakdown of what it does:
func (s *Handlers) InsertUserInfo(w http.ResponseWriter, r *http.Request) {

	// The code snippet you provided is part of a Go program that defines handler functions for processing
	// HTTP requests related to user information.
	var InsertUserInfoReq dto.UserInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&InsertUserInfoReq); err != nil {
		errormessage := errs.ErrorResponse{Errors: []string{strings.ReplaceAll(err.Error(), "json: ", "")}}
		writeResponse(w, http.StatusBadRequest, errormessage)
		return
	}

	// The code snippet you provided is performing input validation on the `InsertUserInfoReq` object using
	// the `ValidateUserInfoRequest` function from the `validator` package. Here's a breakdown of what it
	// does:
	errorResponse := validator.ValidateUserInfoRequest(&InsertUserInfoReq)
	if len(errorResponse) > 0 {
		writeResponse(w, http.StatusBadRequest, errs.ErrorResponse{Errors: errorResponse})
		return
	}

	// The code snippet you provided is calling the `InsertUserInfo` method of the `Service` object stored
	// in the `Handlers` struct. This method is responsible for inserting user information based on the
	// `InsertUserInfoReq` object passed to it.
	UpdateCustomerRes, err := s.Service.InsertUserInfo(&InsertUserInfoReq)
	if err != nil {
		errormessage := errs.ErrorResponse{Errors: err.Errors}
		writeResponse(w, err.StatusCode, errormessage)
		return
	}

	writeResponse(w, http.StatusOK, UpdateCustomerRes)
}

// The `AuthUserInfo` function in the code snippet is a handler function for authenticating user
// information. Here's a breakdown of what it does:
func (s *Handlers) AuthUserInfo(w http.ResponseWriter, r *http.Request) {

	// The code snippet you provided is part of a handler function named `AuthUserInfo` that is
	// responsible for authenticating user information.
	var AuthUserInfoRequest dto.AuthUserInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&AuthUserInfoRequest); err != nil {
		errormessage := errs.ErrorResponse{Errors: []string{strings.ReplaceAll(err.Error(), "json: ", "")}}
		writeResponse(w, http.StatusBadRequest, errormessage)
		return
	}

	// The code snippet `UpdateCustomerRes, err := s.Service.AuthUserInfo(&AuthUserInfoRequest)` is
	// calling the `AuthUserInfo` method of the `Service` object stored in the `Handlers` struct. This
	// method is responsible for authenticating user information based on the `AuthUserInfoRequest` object
	// passed to it.
	UpdateCustomerRes, err := s.Service.AuthUserInfo(&AuthUserInfoRequest)
	if err != nil {
		errormessage := errs.ErrorResponse{Errors: err.Errors}
		writeResponse(w, err.StatusCode, errormessage)
		return
	}

	writeResponse(w, http.StatusOK, UpdateCustomerRes)
}

// The `writeResponse` function sets the necessary headers and writes a JSON response with the provided
// data and status code.
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
