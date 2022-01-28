package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vanshajg/go-play/container"
	"github.com/vanshajg/go-play/models/dto"
	"github.com/vanshajg/go-play/service"
)

type CommentController struct {
	container container.Container
	service   *service.CommentService
}

func NewCommentController(c container.Container) *CommentController {
	return &CommentController{container: c, service: service.NewCommentService(c)}
}

// CreateComment create a new comment by http post.
// @Summary Create a new Comment
// @Description Create a new Comment
// @Tags Comments
// @Accept  json
// @Produce  json
// @Param data body dto.CommentDto true "a new comment"
// @Success 200 {object} models.Comment "Successfully created a comment"
// @Failure 400 {string} message "Failed to create."
// @Failure 401 {boolean} bool "Failed to the authenticate. Returns false."
// @Router /api/comment [post]
func (ctrl *CommentController) CreateComment(c echo.Context) error {
	dto := dto.NewCommentDto()
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, dto)
	}
	result, status := ctrl.service.CreateComment(dto)
	if status != nil {
		return c.JSON(http.StatusBadRequest, status)
	}
	return c.JSON(http.StatusOK, result)
}

// GetComment gets comment list
// @Summary gets comment list
// @Description
// @Tags Comments
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Comment
// @Failure 400 {string} message "failed to get comments"
// @Failure 401 {boolean} bool "Failed to the authenticate. Returns false."
// @Router /api/comment [get]
func (ctrl *CommentController) GetComment(c echo.Context) error {
	page, err := ctrl.service.GetComments()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, page)
}
