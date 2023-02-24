package categoryrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

// ListCategoryStore interface
type ListCategoryStore interface {
	ListCategoryByCondition(ctx context.Context, conditions map[string]interface{}) ([]categorymodel.Category, error)
}

// listCategoryRepo interface interact with storage layer
type listCategoryRepo struct {
	store ListCategoryStore
}

// NewListCategoryRepo func create new listCategoryRepo
func NewListCategoryRepo(store ListCategoryStore) *listCategoryRepo {
	return &listCategoryRepo{store: store}
}

// ListCategory func
func (r listCategoryRepo) ListCategory(ctx context.Context) ([]categorymodel.Category, error) {
	result, err := r.store.ListCategoryByCondition(ctx, nil)

	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return result, nil
}
