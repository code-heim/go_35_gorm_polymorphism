package controllers

import (
	"fmt"
	"gin_blogs/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PhotosIndex(c *gin.Context) {
	photos := models.PhotosAll(c)
	c.HTML(
		http.StatusOK,
		"photos/index.tpl",
		gin.H{
			"photos":     photos,
			"page":       c.GetInt("page"),
			"pageSize":   c.GetInt("pageSize"),
			"totalPages": c.GetInt("totalPages"),
			"UserID":     c.GetUint("userID"),
		},
	)
}

func PhotosShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	photo := models.PhotosFind(id)
	c.HTML(
		http.StatusOK,
		"photos/show.tpl",
		gin.H{
			"photo":  photo,
			"UserID": c.GetUint("userID"),
		},
	)
}
