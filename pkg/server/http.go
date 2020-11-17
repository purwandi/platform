package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Component ...
type Component interface {
	Mount(*echo.Echo)
}

// Server ...
type Server struct {
	context context.Context
	router  *echo.Echo
}

// NewServer ...
func NewServer(ctx context.Context) *Server {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	return &Server{
		context: ctx,
		router:  e,
	}
}

// Register is fo registering router component
func (s *Server) Register(route Component) {
	route.Mount(s.router)
}

// Serve is for serve router
func (s *Server) Serve() {
	go func() {
		if err := s.router.Start(":8080"); err != nil {
			s.router.Logger.Fatal("shutting down the server")
		}
	}()
}

// Close for closing the server
func (s *Server) Close() error {
	return s.router.Shutdown(s.context)
}
