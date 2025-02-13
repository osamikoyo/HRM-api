package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/config"
	"github.com/osamikoyo/hrm-api/internal/vocation"
	"github.com/osamikoyo/hrm-api/internal/worker"
	"github.com/osamikoyo/hrm-api/pkg/loger"
)

type Server struct{
	Echo *echo.Echo
	Config *config.Config
	Loger loger.Logger
}

func New() *Server {
	return &Server{
		Echo: echo.New(),
		Config: config.LoadConfig(),
		Loger: loger.New(),
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func ()  {
		<- ctx.Done()
		s.Loger.Info().Msg("Server shutdown!:3")
		s.Echo.Shutdown(ctx)	
	}()

	s.Loger.Info().Msg("Registing routers...")

	worker.RegisterRoutes(s.Echo, s.Config)
	vocation.RegisterRoutes(s.Echo, s.Config)

	s.Loger.Info().Msg("Starting http server...")

	return s.Echo.Start(fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
}