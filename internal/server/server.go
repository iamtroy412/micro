package server

import (
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
}
