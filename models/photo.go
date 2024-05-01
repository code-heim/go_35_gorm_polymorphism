package models

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Url         string    `gorm:"size:255"`
	Description string    `gorm:"type:text"`
	Comments    []Comment `gorm:"polymorphic:Commentable;"`
}

func PhotosAll(ctx *gin.Context) *[]Photo {
	var photos []Photo
	DB.Preload("Comments").Where("deleted_at is NULL").Scopes(Paginate(ctx)).Order("updated_at desc").Find(&photos)
	fmt.Println(photos)
	return &photos
}

func PhotosFind(id uint64) *Photo {
	var photo Photo
	DB.Preload("Comments").Where("id = ?", id).First(&photo)
	return &photo
}
