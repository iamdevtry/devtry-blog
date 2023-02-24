package categoryrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

// CreateCategoryStorage interface to interact with storage
type CreateCategoryStore interface {
	Create(ctx context.Context, data *categorymodel.CategoryCreate) error
	Find(ctx context.Context, cond map[string]interface{}) (*categorymodel.Category, error)
}

// createCategoryRepo struct
type createCategoryRepo struct {
	store CreateCategoryStore
}

// NewCreateCategoryRepo create new repo
func NewCreateCategoryRepo(store CreateCategoryStore) *createCategoryRepo {
	return &createCategoryRepo{store: store}
}

// CreateCategory create new category
func (r *createCategoryRepo) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	if data.ParentId != 0 {
		parentCat, err := r.store.Find(ctx, map[string]interface{}{"id": data.ParentId})

		if err != nil {
			return common.ErrCannotGetEntity(categorymodel.EntityName, fmt.Errorf("cannot find parent category with id %d", data.ParentId))
		}

		if parentCat.Status == 0 {
			return common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("parent category not found"))
		}
	}

	if err := r.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
