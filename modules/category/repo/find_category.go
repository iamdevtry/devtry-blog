package categoryrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

type FindCategoryStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*categorymodel.Category, error)
}

type findCategoryRepo struct {
	store FindCategoryStore
}

func NewFindCategoryRepo(store FindCategoryStore) *findCategoryRepo {
	return &findCategoryRepo{store: store}
}

func (r *findCategoryRepo) FindCategory(ctx context.Context, cond map[string]interface{}) (*categorymodel.Category, error) {
	data, err := r.store.Find(ctx, cond)
	if err != nil {
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	return data, nil
}
