package db

import (
	"github.com/mercadolibre/store_oauth-api/src/app/clients/cassandra"
	"github.com/mercadolibre/store_oauth-api/src/domain/access_token"
	"github.com/mercadolibre/store_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()

	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
