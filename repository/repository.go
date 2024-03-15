package repo

import (
	"net/http"
	"sangamCyber/database"
	"sangamCyber/dto"
	"sangamCyber/entity"
	"sangamCyber/errs"

	"github.com/jmoiron/sqlx"
)

type GetDbClient func() *sqlx.DB

type RepositoryDb struct {
	client GetDbClient
}

type Respository interface {
	ValidateUser(*dto.UserInfoRequest) (entity.UserInfoResponse, *errs.AppError)
	InsertUserInfoReq(*dto.UserInfoRequest) (*dto.InsertUserResponse, *errs.AppError)
	AuthUserInfo(*dto.AuthUserInfoRequest) (*entity.AuthEntityResponse, *errs.AppError)
}

// This function `ValidateUser` in the `RepositoryDb` struct is responsible for validating user
// information based on the provided `UserInfoRequest`.
func (r RepositoryDb) ValidateUser(req *dto.UserInfoRequest) (entity.UserInfoResponse, *errs.AppError) {
	var errorResponse []string
	client, DbError := database.GetDbClient()

	if DbError != nil {
		return entity.UserInfoResponse{}, errs.NewUnexpectedError()
	}

	queryUserName := `SELECT user_id
	FROM user_data
	WHERE username = $1`

	queryEmailId := `SELECT user_id
	FROM user_data
	WHERE email_id = $1`

	var userNameResponse []entity.UserID
	var emailIdResponse []entity.UserID

	userNameErr := client.Select(&userNameResponse, queryUserName,
		req.UserName,
	)

	emailIdErr := client.Select(&emailIdResponse, queryEmailId,
		req.EmailID,
	)

	var response entity.UserInfoResponse
	response.UsernameInfo = userNameResponse
	response.EmailInfo = emailIdResponse

	if userNameErr != nil || emailIdErr != nil {
		client.Close()
		errorResponse = append(errorResponse, userNameErr.Error())
		return entity.UserInfoResponse{}, errs.ValidateResponse(errorResponse, http.StatusInternalServerError)
	}

	return response, nil
}

// The `InsertUserInfoReq` function in the `RepositoryDb` struct is responsible for inserting user
// information into the database based on the provided `UserInfoRequest`.
func (r RepositoryDb) InsertUserInfoReq(req *dto.UserInfoRequest) (*dto.InsertUserResponse, *errs.AppError) {
	var errorResponse []string
	client, DbError := database.GetDbClient()

	if DbError != nil {
		return nil, errs.NewUnexpectedError()
	}

	query := `INSERT INTO user_data (username, password, email_id, created_by, modified_by)
	VALUES($1, $2, $3,'admin','admin') RETURNING user_id as userId`
	var response int
	execErr := client.QueryRow(query,
		req.UserName,
		req.Password,
		req.EmailID,
	).Scan(&response)

	if execErr != nil {
		client.Close()
		errorResponse = append(errorResponse, execErr.Error())
		return nil, errs.ValidateResponse(errorResponse, http.StatusInternalServerError)
	}

	return &dto.InsertUserResponse{
		UserId: response,
	}, nil
}

// The `AuthUserInfo` function in the `RepositoryDb` struct is responsible for authenticating user
// information based on the provided `AuthUserInfoRequest`.
func (r RepositoryDb) AuthUserInfo(req *dto.AuthUserInfoRequest) (*entity.AuthEntityResponse, *errs.AppError) {
	var errorResponse []string
	client, DbError := database.GetDbClient()

	if DbError != nil {
		return nil, errs.NewUnexpectedError()
	}

	query := `SELECT user_id,password from user_data WHERE username=$1`
	response := make([]entity.AuthEntityResponse, 0)
	execErr := client.Select(&response, query,
		req.UserName,
	)

	if execErr != nil {
		client.Close()
		errorResponse = append(errorResponse, execErr.Error())
		return nil, errs.ValidateResponse(errorResponse, http.StatusInternalServerError)
	}

	if len(response) == 0 {
		return nil, errs.ValidateResponse([]string{"Invalid userName"}, http.StatusBadRequest)
	}

	return &entity.AuthEntityResponse{
		Password: response[0].Password,
		UserID:   response[0].UserID,
	}, nil
}

func NewRepositoryDb() RepositoryDb {
	return RepositoryDb{}
}
