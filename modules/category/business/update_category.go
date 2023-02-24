package categorybusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

type UpdateCategoryRepo interface {
	UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error
}

type updateCategoryBusiness struct {
	repo UpdateCategoryRepo
}

func NewUpdateCategoryBusiness(repo UpdateCategoryRepo) *updateCategoryBusiness {
	return &updateCategoryBusiness{repo: repo}
}

func (b *updateCategoryBusiness) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	if err := b.repo.UpdateCategory(ctx, id, data); err != nil {
		return common.ErrCannotUpdatedEntity(categorymodel.EntityName, err)
	}

	return nil
}
