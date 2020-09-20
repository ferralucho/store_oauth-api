package app

import (
	"github.com/ferralucho/store_oauth-api/src/domain/access_token"
	"github.com/ferralucho/store_oauth-api/src/http"
	"github.com/ferralucho/store_oauth-api/src/repository/db"
	"github.com/ferralucho/store_oauth-api/src/repository/rest"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(access_token.NewService(db.NewRepository(), rest.NewRestUsersRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8083")
}
