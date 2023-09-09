package server

import (
	"log"
	"net/http"

	"github.com/iamtroy412/micro/internal/database"
	"github.com/labstack/echo/v4"
)

type Server interface {
    Start() error
}

type EchoServer struct {
    echo *echo.Echo
    DB database.DatabseClient
}

func NewEchoServer(db database.DatabseClient) Server {
    server := &EchoServer{
        echo: echo.New(),
        DB: db,
    }
    server.registerRoutes()
    return server
}

func (s *EchoServer) Start() error {
    if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server shutdown occurred: %s", err)
        return err
    }
    return nil
}

func (s *EchoServer) registerRoutes() {

}
