package http

import (
	"net"
	"net/http"
	"strings"

	"github.com/kyonko-moe/kagami/model"
	mhttp "github.com/kyonko-moe/kagami/model/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Register(c *gin.Context) {
	var (
		err error
		arg = &mhttp.ArgRegister{}
		n   *model.Node
		ip  = net.ParseIP(strings.Split(c.Request.RemoteAddr, ":")[0])
	)
	if err = bind(c, arg); err != nil {
		return
	}
	if ip == nil {
		ip = net.ParseIP("127.0.0.1")
	}
	if n, err = s.tracker.Register(c, ip, arg.Name); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := &mhttp.RespRegister{
		ID:   n.ID,
		Name: n.Name,
		Type: string(n.Type),
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) Locate(c *gin.Context) {
	var (
		err error
		arg = &mhttp.ArgLocate{}
		n   *model.Node
	)
	if err = bind(c, arg); err != nil {
		return
	}
	if n, err = s.tracker.Locate(c, arg.Name); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := &mhttp.RespLocate{
		ID:   n.ID,
		Name: n.Name,
		Type: string(n.Type),
		IPv4: n.IP.String(),
	}
	c.JSON(http.StatusOK, resp)
}

func (s *Server) Connect(c *gin.Context) {
	var (
		err  error
		arg  = &mhttp.ArgConnect{}
		addr *net.UDPAddr
	)
	if err = bind(c, arg); err != nil {
		return
	}
	if addr, err = s.udpPuncher.Connect(c, net.ParseIP(arg.IPv4)); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp := &mhttp.RespConnect{
		IPv4: addr.IP.String(),
		Port: addr.Port,
	}
	c.JSON(http.StatusOK, resp)
}
