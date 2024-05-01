package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title    string    `gorm:"size:255"`
	Content  string    `gorm:"type:text"`
	Comments []Comment `gorm:"polymorphic:Commentable;"`
}

func BlogsAll(ctx *gin.Context) *[]Blog {
	var blogs []Blog
	DB.Preload("Comments").Where("deleted_at is NULL").Scopes(Paginate(ctx)).Order("updated_at desc").Find(&blogs)
	fmt.Println(blogs)
	return &blogs
}

func BlogsFind(id uint64) *Blog {
	var blog Blog
	DB.Preload("Comments").Where("id = ?", id).First(&blog)
	return &blog
}
