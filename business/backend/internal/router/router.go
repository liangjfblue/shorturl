package router

import (
	"backend/internal/config"
	"backend/internal/handler"
	"common/constinfo"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewRouter)

func NewRouter(
	c *config.Config,
	shortHandler *handler.ShortHandler,
) *gin.Engine {
	if c.Server.Env == constinfo.EnvPro {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

	router.LoadHTMLGlob("templates/*")

	apiGroup := router.Group("")

	initShortRouter(apiGroup, shortHandler)

	return router
}
