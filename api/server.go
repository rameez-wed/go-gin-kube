package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/go-gin-kube/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(dbStore db.Store) *Server {
	server := &Server{store: dbStore}
	router := gin.Default()
	// router.GET("/posts", server.getPosts)
	router.GET("/authors/:id", server.getAuthor)
	router.GET("/authors", server.getAuthors)
	router.POST("/authors", server.createAuthor)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
