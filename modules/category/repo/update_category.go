package categoryrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

type UpdateCategoryStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*categorymodel.Category, error)
	Update(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error
}

type updateCategoryRepo struct {
	store UpdateCategoryStore
}

func NewUpdateCategoryRepo(store UpdateCategoryStore) *updateCategoryRepo {
	return &updateCategoryRepo{store: store}
}

func (repo *updateCategoryRepo) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	oldData, err := repo.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("category not found"))
	}

	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	if data.ParentId != 0 {
		parentCat, err := repo.store.Find(ctx, map[string]interface{}{"id": data.ParentId})

		if err != nil {
			return common.ErrCannotGetEntity(categorymodel.EntityName, fmt.Errorf("cannot find parent category with id %d", data.ParentId))
		}

		if parentCat.Status == 0 {
			return common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("parent category not found"))
		}
	}

	if err := repo.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdatedEntity(categorymodel.EntityName, err)
	}

	return nil
}
