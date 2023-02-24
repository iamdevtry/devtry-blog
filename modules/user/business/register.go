package userbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	usermodel "github.com/iamdevtry/blog/modules/user/model"
)

type RegisterRepo interface {
	Register(ctx context.Context, data *usermodel.UserCreate) error
}

type registerBusiness struct {
	repo RegisterRepo
}

func NewRegisterBusiness(repo RegisterRepo) *registerBusiness {
	return &registerBusiness{repo: repo}
}

func (b *registerBusiness) RegisterUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := b.repo.Register(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
