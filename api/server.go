package api

import (
	"fmt"
	"log"

	db "github.com/Kcih4518/simpleBank_2023/db/sqlc"
	"github.com/Kcih4518/simpleBank_2023/token"
	"github.com/Kcih4518/simpleBank_2023/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// register various handler functions
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("currency", validCurrency); err != nil {
			// Handle the error, for example, log it or return an error
			log.Printf("Failed to register 'currency' custom validation: %v", err)
			// You can take appropriate error-handling measures here, such as returning an error message
		}
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// user
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// account
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.DELETE("/accounts/:id", server.delAccount)
	router.GET("/accounts", server.listAccount)
	router.PATCH("/accounts/:id", server.updateAccount)

	// transfer
	router.POST("/transfers", server.createTransfer)

	server.router = router
}

// H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
