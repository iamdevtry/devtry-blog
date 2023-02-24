package uploadmodel

import (
	"errors"

	"github.com/iamdevtry/blog/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline`
}

func (Upload) TableName() string {
	return "images"
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}

func ErrCannotDeleteFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot delete file",
		"ErrCannotDeleteFile",
	)
}
