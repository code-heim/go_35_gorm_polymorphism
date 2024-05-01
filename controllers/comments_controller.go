package controllers

import (
	"gin_blogs/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CommentsList(c *gin.Context) {
	// The URL can be blogs/<id>/comments or photos/<id>/comments
	commentableType, commentableID := readCommentable(c)
	// Get comments from the model
	comments := models.CommentsFind(commentableType, commentableID)
	c.HTML(
		http.StatusOK,
		"comments/list.tpl",
		gin.H{
			"comments": comments,
			"UserID":   c.GetUint("userID"),
		},
	)
}

func CommentsCreate(c *gin.Context) {
	commentableType, commentableID := readCommentable(c)
	type commentFormData struct {
		Comment string `form:"comment"`
	}
	var data commentFormData
	c.Bind(&data)
	models.CommentsCreate(commentableType, commentableID, data.Comment)
	c.Redirect(http.StatusFound, c.Request.RequestURI)
}

func readCommentable(c *gin.Context) (string, uint64) {
	splits := strings.Split(c.Request.RequestURI, "/")
	commentableType := strings.ToLower(splits[1])
	commentableID, _ := strconv.Atoi(splits[2])
	return commentableType, uint64(commentableID)
}
