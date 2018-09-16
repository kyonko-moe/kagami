package http

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/kyonko-moe/kagami/biz/track"
	"github.com/kyonko-moe/kagami/model"
	mhttp "github.com/kyonko-moe/kagami/model/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Server struct {
	r     *gin.Engine
	track *track.Tracker
}

func New() *Server {
	return &Server{
		track: track.New(),
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

func (s *Server) Register(c *gin.Context) {
	var (
		err error
		p   = &mhttp.ParamRegister{}
		n   *model.Node
		ip  = net.ParseIP(strings.Split(c.Request.RemoteAddr, ":")[0])
	)
	if err = c.Bind(p); err != nil {
		err = errors.WithStack(err)
		log.Fatalf("%+v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	ipAddr := &net.IPAddr{IP: ip}
	if n, err = s.track.Register(c, ipAddr, p.Name); err != nil {
		log.Printf("%+v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, n)
}

func (s *Server) Locate(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (s *Server) Connect(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
