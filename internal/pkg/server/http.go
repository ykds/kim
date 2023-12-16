package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"kim/internal/pkg/log"
	"net/http"
	"time"
)

type HttpServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

type Server struct {
	log        *log.Logger
	httpServer *http.Server
	engine     *gin.Engine
}

func NewHttpServer(cfg HttpServerConfig, logger *log.Logger) *Server {
	engine := gin.New()
	engine.Use(gin.RecoveryWithWriter(logger.GetOut()), gin.LoggerWithWriter(logger.GetOut()))
	gin.SetMode(gin.ReleaseMode)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: engine,
	}
	return &Server{
		log:        logger,
		httpServer: server,
		engine:     engine,
	}
}

func (s *Server) RegisterRouter(fns ...func(r *gin.Engine)) {
	for _, fn := range fns {
		fn(s.engine)
	}
}

func (s *Server) Run() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		s.log.Errorf("http server down, err: %v", err)
	}
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.log.Errorf("http server shutdown err: %v", err)
	}
}
