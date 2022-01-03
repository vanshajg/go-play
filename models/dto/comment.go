package dto

import "github.com/vanshajg/go-play/models"

type CommentDto struct {
	By     string `json:"by"`
	ID     uint   `json:"id"`
	Parent uint   `json:"parent"`
	Text   string `json:"text"`
	Time   uint   `json:"time"`
	Type   string `json:"type"`
}

func NewCommentDto() *CommentDto {
	return &CommentDto{}
}

func (c *CommentDto) Create() *models.Comment {
	return models.NewComment()
}
