package categorybusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

// ListCategoryRepo interface
type ListCategoryRepo interface {
	ListCategory(ctx context.Context) ([]categorymodel.Category, error)
}

// listCategoryBusiness interface interact with repo layer
type listCategoryBusiness struct {
	repo ListCategoryRepo
}

// NewListCategoryBusiness func create new listCategoryBusiness
func NewListCategoryBusiness(repo ListCategoryRepo) *listCategoryBusiness {
	return &listCategoryBusiness{repo: repo}
}

// ListCategory func
func (b listCategoryBusiness) ListCategory(ctx context.Context) ([]categorymodel.Category, error) {
	result, err := b.repo.ListCategory(ctx)

	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return result, nil
}
