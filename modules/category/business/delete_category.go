package categorybusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

type DeleteCategoryRepo interface {
	SoftDeleteCategory(ctx context.Context, id string) error
}

type deleteCategoryBusiness struct {
	repo DeleteCategoryRepo
}

func NewDeleteCategoryBusiness(repo DeleteCategoryRepo) *deleteCategoryBusiness {
	return &deleteCategoryBusiness{repo: repo}
}

func (b *deleteCategoryBusiness) DeleteCategory(ctx context.Context, id string) error {
	if err := b.repo.SoftDeleteCategory(ctx, id); err != nil {
		return common.ErrCannotDeletedEntity(categorymodel.EntityName, err)
	}

	return nil
}
