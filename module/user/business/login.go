package userbusiness

import (
	"context"
	"gitlab.com/genson1808/food-delivery/common"
	"gitlab.com/genson1808/food-delivery/component/tokenprovider"
	usermodel "gitlab.com/genson1808/food-delivery/module/user/model"
)

type LoginStorage interface {
	FindWithCondition(ctx context.Context, conditions map[string]any, moreKeys ...string) (*usermodel.User, error)
}

type LoginBiz struct {
	store         LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(store LoginStorage, hasher Hasher, tokenprovider tokenprovider.Provider, expiry int) *LoginBiz {
	return &LoginBiz{store: store, tokenProvider: tokenprovider, hasher: hasher, expiry: expiry}
}

func (biz *LoginBiz) Login(ctx context.Context, data *usermodel.UserCredentials) (*tokenprovider.Token, error) {
	user, err := biz.store.FindWithCondition(ctx, map[string]any{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payLoad := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payLoad, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return accessToken, nil
}
