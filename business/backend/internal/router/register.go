package router

import (
	"backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func initShortRouter(group *gin.RouterGroup, h *handler.ShortHandler) {
	group.GET("/:url", h.Short)
}
