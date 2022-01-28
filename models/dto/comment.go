package dto

import (
	"github.com/vanshajg/go-play/models"
)

type CommentDto struct {
	By     string `json:"by"`
	ID     uint   `json:"id"`
	Parent uint   `json:"parent"`
	Kids   []int  `json:"kids"`
	Text   string `json:"text"`
	Time   uint   `json:"time"`
	Type   string `json:"type"`
}

func NewCommentDto() *CommentDto {
	return &CommentDto{}
}

func (c *CommentDto) Create() *models.Comment {
	return models.NewComment(
		c.ID,
		c.Text,
		isRemote(c.Text),
		[]int{},
		[]int{},
	)
}

func isRemote(text string) bool {
	//  regexp.Match("remote.{0,4}true", []byte(text))

	// todo regex match with remote : true
	return true
}
