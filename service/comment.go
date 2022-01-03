package service

import (
	"github.com/vanshajg/go-play/container"
	"github.com/vanshajg/go-play/models"
	"github.com/vanshajg/go-play/models/dto"
	"github.com/vanshajg/go-play/repository"
)

type CommentService struct {
	container container.Container
}

func NewCommentService(c container.Container) *CommentService {
	return &CommentService{container: c}
}

func (c *CommentService) CreateComment(dto *dto.CommentDto) (*models.Comment, map[string]string) {
	rep := c.container.GetRepository()
	var result *models.Comment
	err := rep.Transaction(func(r repository.Repository) error {
		comment := dto.Create()
		var err error
		if result, err = comment.Create(r); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, map[string]string{"error": "failed to create comment"}
	}
	return result, nil
}

func (c *CommentService) GetComments() (*[]models.Comment, error) {
	rep := c.container.GetRepository()
	comment := models.Comment{}
	result, err := comment.GetAll(rep)
	if err != nil {
		c.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil, err
	}

	return result, nil

}
