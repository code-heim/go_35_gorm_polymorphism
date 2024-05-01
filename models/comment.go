package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentableType string
	CommentableID   uint64 `gorm:"size:64"`
	Content         string `gorm:"type:text"`
}

func CommentsFind(commentableType string, commentableID uint64) *[]Comment {
	var comments []Comment
	DB.Where("commentable_type = ? and commentable_id = ?", commentableType, commentableID).Find(&comments)
	return &comments
}

func CommentsCreate(commentableType string, commentableID uint64, content string) *Comment {
	entry := Comment{CommentableType: commentableType,
		CommentableID: commentableID,
		Content:       content}
	DB.Create(&entry)
	return &entry
}
