package models

import (
	"database/sql"

	"github.com/vanshajg/go-play/repository"
	"github.com/vanshajg/go-play/utils"
	"gorm.io/gorm"
)

type Comment struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Text     string `json:"text"`
	IsRemote bool   `json:"is_remote"`
}

// TableName returns the table name of book struct and it is used by gorm.
func (Comment) TableName() string {
	return "comment"
}

const (
	selectComment = "select * from comment"
)

func NewComment(id uint, text string, isRemote bool, tech []int, locations []int) *Comment {
	return &Comment{
		ID:       id,
		Text:     text,
		IsRemote: isRemote,
	}
}

func (c *Comment) Create(rep repository.Repository) (*Comment, error) {
	if err := rep.Select("id", "is_remote", "text").Create(c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Comment) GetAll(rep repository.Repository) (*[]Comment, error) {
	return findRows(rep, selectComment, "", "", []interface{}{})
}

func findRows(rep repository.Repository, sqlquery string, page string, size string, args []interface{}) (*[]Comment, error) {
	var comments []Comment
	var rows *sql.Rows
	var rec Comment
	var err error
	if rows, err = createRaw(rep, sqlquery, page, size, args).Rows(); err != nil {
		return nil, err
	}
	for rows.Next() {
		if err = rep.ScanRows(rows, &rec); err != nil {
			return nil, err
		}
		comments = append(comments, rec)
	}
	return &comments, nil
}

func createRaw(rep repository.Repository, sql string, pageNum string, pageSize string, args []interface{}) *gorm.DB {
	if utils.IsNum(pageNum) && utils.IsNum(pageSize) {
		page := utils.ConvertToInt(pageNum)
		size := utils.ConvertToInt(pageSize)
		args = append(args, size, page*size)
		sql += "limit ? offset ?"
	}
	if len(args) > 0 {
		return rep.Raw(sql, args...)
	}
	return rep.Raw(sql)
}
