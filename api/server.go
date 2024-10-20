package api

import (
	db "github.com/gaberingo/SimpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server amd setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.getListAccount)
	router.GET("/transfers", server.getListTransfer)

	// Add routes to router
	server.router = router

	return server
}

// start run the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
