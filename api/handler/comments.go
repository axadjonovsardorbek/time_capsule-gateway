package handler

import (
	"context"
	"fmt"
	cp "gateway/genproto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CommentCreate handles the creation of a new comment.
// @Summary Create comment
// @Description Create a new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body cp.CommentsCreateReq true "Comment data"
// @Success 200 {object} string "Comment created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /memory/{id}/comment [post]
func (h *Handler) CommentCreate(c *gin.Context) {
	var req cp.CommentsCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.Comment.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created"})
}

// CommentGetById handles the get a comment.
// @Summary Get comment
// @Description Get a comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} cp.CommentsGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /memory/{id}/comment/{id} [get]
func (h *Handler) CommentGetById(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	res, err := h.srvs.Comment.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get comment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CommentGetAll handles getting all comment.
// @Summary Get all comment
// @Description Get all comment
// @Tags comment
// @Accept json
// @Produce json
// @Param user_id query string false "UserId"
// @Param memory_id query string false "MemoryId"
// @Param page query integer false "Page"
// @Success 200 {object} cp.CommentsGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /memory/{id}/comment/all [get]
func (h *Handler) CommentGetAll(c *gin.Context) {
	req := cp.CommentsGetAllReq{
		UserId:   c.Query("user_id"),
		MemoryId: c.Query("memory_id"),
		Filter:   &cp.Filter{},
	}

	pageStr := c.Query("page")
	var page int
	var err error
	if pageStr == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
			return
		}
	}

	filter := cp.Filter{
		Page: int32(page),
	}

	req.Filter.Page = filter.Page

	res, err := h.srvs.Comment.GetAll(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get comments", "details": err.Error()})
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

// CommentUpdate handles updating an existing comment.
// @Summary Update comment
// @Description Update an existing comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string false "Id"
// @Param content query string false "Content"
// @Success 200 {object} string "Comment updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Comment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /memory/{id}/comment/{id} [put]
func (h *Handler) CommentUpdate(c *gin.Context) {
	memory := cp.CommentsUpdateReq{
		Id:      c.Query("id"),
		Content: c.Query("content"),
	}

	_, err := h.srvs.Comment.Update(context.Background(), &memory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update comment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated"})
}

// CommentDelete handles deleting a comment by ID.
// @Summary Delete comment
// @Description Delete a comment by ID
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} string "Comment deleted"
// @Failure 400 {object} string "Invalid comment ID"
// @Failure 404 {object} string "Comment not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /memory/{id}/comment/{id} [delete]
func (h *Handler) CommentDelete(c *gin.Context) {
	id := &cp.ById{Id: c.Param("id")}
	_, err := h.srvs.Comment.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete comment", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
