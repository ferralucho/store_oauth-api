package access_token

import (
	"strings"

	"github.com/mercadolibre/store_oauth-api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(accessTokenID string) (*AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetByID(accessTokenID)

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(AccessToken) *errors.RestErr {
	return nil
}

func (s *service) UpdateExpirationTime(AccessToken) *errors.RestErr {
	return nil
}
