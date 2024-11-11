package handler

import (
	"backend/internal/service"
	"common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShortHandler struct {
	svcShort *service.SvcShort
}

func NewShortHandler(svcShort *service.SvcShort) *ShortHandler {
	r := &ShortHandler{
		svcShort: svcShort,
	}

	return r
}

// Short 短连接访问
func (h *ShortHandler) Short(c *gin.Context) {
	var (
		rsp utils.Result
		ctx = c.Request.Context()
	)

	shortUrl := c.Param("url")
	if len(shortUrl) == 0 || len(shortUrl) > 7 {
		rsp.BadRequest("param failed")
		c.JSON(http.StatusOK, &rsp)
		return
	}

	logUrl, err := h.svcShort.GetLongUrl(ctx, shortUrl)
	if err != nil {
		rsp.InternalServerError("系统错误")
		c.JSON(http.StatusOK, &rsp)
		return
	}
	if logUrl == "" {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}

	c.Redirect(http.StatusFound, logUrl)
}
