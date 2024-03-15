package service

import (
	"net/http"
	"sangamCyber/converter"
	"sangamCyber/dto"
	"sangamCyber/errs"
	domain "sangamCyber/repository"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo domain.Respository
}

type Service interface {
	InsertUserInfo(req *dto.UserInfoRequest) (*dto.InsertUserResponse, *errs.AppError)
	AuthUserInfo(req *dto.AuthUserInfoRequest) (*dto.AuthUserResponse, *errs.AppError)
}

// This `InsertUserInfo` function is responsible for inserting user information into the system.
func (r service) InsertUserInfo(req *dto.UserInfoRequest) (*dto.InsertUserResponse, *errs.AppError) {

	var errors []string

	validateResponse, Err := r.repo.ValidateUser(req)

	// This block of code in the `InsertUserInfo` function is checking if the email or username provided
	// in the request is already registered in the system.
	if len(validateResponse.EmailInfo) > 0 {
		errors = append(errors, "This email is already registered")
	}

	if len(validateResponse.UsernameInfo) > 0 {
		errors = append(errors, "This userName is already registered")
	}

	if len(errors) > 0 {
		return nil, errs.ValidateResponse(errors, http.StatusBadRequest)
	}

	if Err != nil {
		return nil, Err
	}

	// The code snippet you provided is responsible for hashing the user's password before inserting it
	// into the system.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	// The code snippet you provided is iterating over each byte in the hashed password byte slice and
	// converting each byte from English to Tamil using the `converter.EnglishToTamil` function. It then
	// concatenates the converted Tamil characters with a space character and stores the result in the
	// `hashTamil` variable. This process essentially converts the hashed password from English characters
	// to Tamil characters for some specific purpose in your application.
	var hashTamil string
	for _, char := range hashedPassword {
		hashTamil += converter.EnglishToTamil(byte(char)) + " "
	}

	req.Password = hashTamil
	response, Err := r.repo.InsertUserInfoReq(req)

	if Err != nil {
		return nil, Err
	}

	return response, nil
}

// This `AuthUserInfo` function is responsible for authenticating user information.
func (r service) AuthUserInfo(req *dto.AuthUserInfoRequest) (*dto.AuthUserResponse, *errs.AppError) {
	// The code snippet `response, Err := r.repo.AuthUserInfo(req)` is calling the `AuthUserInfo` method
	// on the repository interface to authenticate user information based on the provided request. It
	// retrieves the response and an error (if any) from the repository method call.
	response, Err := r.repo.AuthUserInfo(req)
	if Err != nil {
		return nil, Err
	}

	// The code snippet you provided is part of the `AuthUserInfo` function and it is responsible for
	// converting the hashed password stored in Tamil characters back to English characters for comparison
	// during user authentication.
	var hashEnglish string
	for _, char := range strings.Split(response.Password, " ") {
		englishChar := converter.TamilToEnglish(string(char))
		hashEnglish += englishChar
	}

	// The code snippet you provided is part of the `AuthUserInfo` function and is responsible for
	// authenticating the user's password. Here's a breakdown of what it does:
	err := bcrypt.CompareHashAndPassword([]byte(hashEnglish), []byte(req.Password))
	if err != nil {
		return nil, errs.ValidateResponse([]string{"Invalid Password"}, http.StatusBadRequest)
	} else {
		return &dto.AuthUserResponse{
			UserId: response.UserID,
			Result: "Authentication successful",
		}, nil
	}
}

func NewService(repository domain.Respository) Service {
	return service{repository}
}
