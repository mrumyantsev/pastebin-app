package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpConfig struct {
	Host                string
	Port                int64
	ReadTimeoutSeconds  int64
	WriteTimeoutSeconds int64
	MaxHeaderBytes      int64
	IsEnableDebugMode   bool
}

type HttpServer struct {
	*gin.Engine
	server *http.Server
}

func NewHttpServer(cfg *HttpConfig) *HttpServer {
	if cfg.IsEnableDebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:        engine,
		ReadTimeout:    time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
		MaxHeaderBytes: int(cfg.MaxHeaderBytes),
	}

	return &HttpServer{
		Engine: engine,
		server: server,
	}
}

func (s *HttpServer) Run() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

type AlignedRouterGroup struct {
	*gin.RouterGroup
}

func (s *HttpServer) AlignedGroup(handlers ...gin.HandlerFunc) *AlignedRouterGroup {
	return &AlignedRouterGroup{
		RouterGroup: s.Engine.Group("", handlers...),
	}
}

func (g *AlignedRouterGroup) DELETE__(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return g.DELETE(relativePath, handlers...)
}

func (g *AlignedRouterGroup) GET_____(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return g.GET(relativePath, handlers...)
}

func (g *AlignedRouterGroup) POST____(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return g.POST(relativePath, handlers...)
}

func (g *AlignedRouterGroup) PUT_____(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return g.PUT(relativePath, handlers...)
}

func GetUserId(c *gin.Context) int64 {
	ifaceVal, ok := c.Get("uid")
	if !ok {
		return -1
	}

	val, ok := ifaceVal.(int64)
	if !ok {
		return -1
	}

	return val
}

func SetUserId(c *gin.Context, val int64) {
	c.Set("uid", val)
}

func GetError(c *gin.Context) error {
	ifaceVal, ok := c.Get("err")
	if !ok {
		return nil
	}

	val, ok := ifaceVal.(error)
	if !ok {
		return nil
	}

	return val
}

func SetError(c *gin.Context, val error) {
	c.Set("err", val)
}
