package server

import (
	"github.com/labstack/echo/v4"
	"sf6/internal/store"
)

type Server struct {
	store  *store.Store
	router *echo.Router
}

func NewServer(store *store.Store) (*Server, error) {
	server := &Server{
		store: store,
	}
	return server, nil
}
