package api

import (
	db "github.com/Kcih4518/simpleBank_2023/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// TODO: add routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.DELETE("/accounts/:id", server.delAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
