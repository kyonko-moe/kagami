package http

import (
	"net/http"

	"github.com/kyonko-moe/kagami/biz/punch/udp"
	"github.com/kyonko-moe/kagami/biz/track"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Server struct {
	r          *gin.Engine
	tracker    *track.Tracker
	udpPuncher *udp.Puncher
}

func New() *Server {
	return &Server{
		tracker:    track.New(),
		udpPuncher: udp.New(),
	}
}

func (s *Server) String() string {
	return "http.v1"
}

func (s *Server) Start() error {
	s.r = gin.Default()
	s.r.POST("/register", s.Register)
	s.r.POST("/connect", s.Connect)
	s.r.GET("/locate", s.Locate)

	go func() {
		s.r.Run("0.0.0.0:2333")
	}()
	return nil
}

func (s *Server) Stop() error {
	return nil
}

func bind(ctx *gin.Context, arg interface{}) (err error) {
	if err = ctx.Bind(arg); err != nil {
		err = errors.WithStack(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	return
}
