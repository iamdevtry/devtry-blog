package userrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	usermodel "github.com/iamdevtry/blog/modules/user/model"
)

type RegisterStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerRepo struct {
	store  RegisterStore
	hasher Hasher
}

func NewRegisterRepo(store RegisterStore, hasher Hasher) *registerRepo {
	return &registerRepo{store: store, hasher: hasher}
}

func (r *registerRepo) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := r.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)
	data.Password = r.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := r.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
