package categorybusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

type CreateCategoryRepo interface {
	CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBusiness struct {
	repo CreateCategoryRepo
}

func NewCreateCategoryBusiness(repo CreateCategoryRepo) *createCategoryBusiness {
	return &createCategoryBusiness{repo: repo}
}

func (b *createCategoryBusiness) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := b.repo.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
