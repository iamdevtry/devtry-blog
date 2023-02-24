package postcategorybusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
)

type CreatePostCategoryRepo interface {
	CreatePostCategory(ctx context.Context, data *postcategorymodel.PostCategory) error
}

type createPostCategoryBusiness struct {
	repo CreatePostCategoryRepo
}

func NewCreatePostCategoryBusiness(repo CreatePostCategoryRepo) *createPostCategoryBusiness {
	return &createPostCategoryBusiness{repo: repo}
}

func (b *createPostCategoryBusiness) CreatePostCategory(ctx context.Context, data *postcategorymodel.PostCategory) error {
	if err := b.repo.CreatePostCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postcategorymodel.EntityName, err)
	}
	return nil
}
