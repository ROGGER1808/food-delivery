package userbusiness

import (
	"context"
	"errors"
	"gitlab.com/genson1808/food-delivery/common"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
)

type RegisterStorage interface {
	FindWithCondition(context context.Context, conditions map[string]any, moreKeys ...string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type RegisterBiz struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterBiz(store RegisterStorage, hasher Hasher) *RegisterBiz {
	return &RegisterBiz{store: store, hasher: hasher}
}

func (biz *RegisterBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {

	userFound, _ := biz.store.FindWithCondition(ctx, map[string]any{"email": data.Email})
	if userFound != nil {
		return usermodel.ErrEmailExisted
	}

	if data.Email == "" {
		return errors.New("email is not empty")
	}
	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
