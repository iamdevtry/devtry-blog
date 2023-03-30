package categoryrepo

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

type DeleteCategoryStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*categorymodel.Category, error)
	SoftDelete(ctx context.Context, id string) error
}

type deleteCategoryRepo struct {
	store DeleteCategoryStore
}

func NewDeleteCategoryStore(store DeleteCategoryStore) *deleteCategoryRepo {
	return &deleteCategoryRepo{store: store}
}

func (repo *deleteCategoryRepo) SoftDeleteCategory(ctx context.Context, id string) error {
	oldData, err := repo.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("category not found"))
	}

	if err := repo.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeletedEntity(categorymodel.EntityName, err)
	}

	return nil
}
