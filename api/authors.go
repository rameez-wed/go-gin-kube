package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/go-gin-kube/db/sqlc"
)

type getAuthorRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAuthor(ctx *gin.Context) {
	var req getAuthorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := server.store.GetAuthorById(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, author)
}

type getAuthorsRequest struct {
	Limit  int32 `form:"offset" binding:"required,min=1"`
	Offset int32 `form:"limit" binding:"required,min=5,max=10"`
}

func (server *Server) getAuthors(ctx *gin.Context) {
	var req getAuthorsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	listAuthorParams := db.ListAuthorsParams{Offset: req.Offset, Limit: req.Limit}
	author, err := server.store.ListAuthors(ctx, listAuthorParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, author)
}

type CreateAuthorParams struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createAuthor(ctx *gin.Context) {
	var req CreateAuthorParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := server.store.CreateAuthor(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, author)
}

func (server *Server) defaultRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Service running")
}
