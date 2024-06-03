package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bibarsss/simple-rest-api-go/service/user"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	apiV1 := router.Group("api/v1")

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(apiV1)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
