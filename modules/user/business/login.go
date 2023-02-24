package userbusiness

import (
	"context"

	"github.com/iamdevtry/blog/component/tokenprovider"
	usermodel "github.com/iamdevtry/blog/modules/user/model"
)

type LoginRepo interface {
	Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error)
}

type loginBusiness struct {
	repo LoginRepo
}

func NewLoginBusiness(repo LoginRepo) *loginBusiness {
	return &loginBusiness{repo: repo}
}

func (b *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	token, err := b.repo.Login(ctx, data)
	if err != nil {
		return nil, err
	}
	return token, nil
}
