package userrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component/tokenprovider"
	usermodel "github.com/iamdevtry/blog/modules/user/model"
)

type LoginStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginRepo struct {
	store LoginStore
	// appCtx        component.AppContext
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginRepo(store LoginStore, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginRepo {
	return &loginRepo{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (r *loginRepo) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := r.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := r.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := r.tokenProvider.Generate(payload, r.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// refreshToken, err := r.tokenProvider.Generate(payload, r.tkCfg.GetRtExp())
	// if err != nil {
	// 	return nil, common.ErrInternal(err)
	// }

	// account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
}
