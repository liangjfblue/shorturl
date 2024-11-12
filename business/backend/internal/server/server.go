package server

import (
	"backend/internal/config"
	"common/logger"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
		s.Addr = fmt.Sprintf(":%d", conf.Server.Http.Port)
	} else {
		s.Addr = fmt.Sprintf(":%d", conf.Server.Https.Port)
	}
	return s
}

func (s *Server) Run() error {
	go func() {
		s.l.Info("http server started")
		if s.c.Server.Type == "http" {
			err := s.httpSrv.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		} else {
			err := s.httpSrv.ListenAndServeTLS(s.c.Server.Https.KeyFile, s.c.Server.Https.CertFile)
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	// other service

	return nil
}

func (s *Server) Stop() error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutdown internal ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.l.Info("http server has been stop")
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		return err
	}

	// stop other

	return nil
}
