package db

import (
	"github.com/mercadolibre/store_oauth-api/src/app/clients/cassandra"
	"github.com/mercadolibre/store_oauth-api/src/domain/access_token"
	"github.com/mercadolibre/store_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetByID(id string) (*access_token.AccessToken, *errors.RestErr) {
	_, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	//TODO implement get access token from cassandra
	return nil, errors.NewInternalServerError("database not implemented")
}
