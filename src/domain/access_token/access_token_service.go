package access_token

import (
	"strings"

	"github.com/ferralucho/store_oauth-api/src/repository/rest"
	"github.com/ferralucho/store_utils-go/rest_errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, rest_errors.RestErr)
	Create(AccessToken) rest_errors.RestErr
	UpdateExpirationTime(AccessToken) rest_errors.RestErr
}

type Service interface {
	GetByID(string) (*AccessToken, rest_errors.RestErr)
	Create(request AccessTokenRequest) (*AccessToken, rest_errors.RestErr)
	UpdateExpirationTime(AccessToken) rest_errors.RestErr
}

type service struct {
	repository          Repository
	restUsersRepository rest.RestUsersRepository
}

func NewService(repo Repository, restUsersRepository rest.RestUsersRepository) Service {
	return &service{
		repository:          repo,
		restUsersRepository: restUsersRepository,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, rest_errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, rest_errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetByID(accessTokenID)

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, rest_errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API
	user, err := s.restUsersRepository.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token
	at := GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra
	if err := s.repository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) rest_errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
