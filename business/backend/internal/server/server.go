package server

import (
	"backend/internal/config"
	"common/logger"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	c       *config.Config
	l       *logger.Logger
	httpSrv *http.Server
}

func NewServer(
	c *config.Config,
	l *logger.Logger,
	httpSrv *http.Server,
) *Server {
	return &Server{
		c:       c,
		l:       l,
		httpSrv: httpSrv,
	}
}

func NewHttpServer(
	conf *config.Config,
	router *gin.Engine,
) *http.Server {
	s := &http.Server{
		Handler: router,
	}
	if conf.Server.Type == "http" {
		s.Addr = ":" + fmt.Sprint(conf.Server.Http.Port)
	} else {
		s.Addr = ":" + fmt.Sprint(conf.Server.Https.Port)
	}
	return s
}

func (s *Server) Run() error {
	go func() {
		s.l.Info("http server started")
		if s.c.Server.Type == "http" {
			if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		} else {
		}

	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.l.Info("http server has been stop")
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
