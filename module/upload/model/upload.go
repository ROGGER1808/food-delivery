package uploadmodel

import (
	"errors"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/foundation/fimage"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel
	fimage.Image
}

func (u *Upload) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUpload)
}

func (Upload) TableName() string {
	return "uploads"
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
