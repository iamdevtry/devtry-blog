package categorybusiness

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

type FindCategoryRepo interface {
	FindCategory(ctx context.Context, cond map[string]interface{}) (*categorymodel.Category, error)
}

type findCategoryBusiness struct {
	repo FindCategoryRepo
}

func NewFindCategoryBusiness(repo FindCategoryRepo) *findCategoryBusiness {
	return &findCategoryBusiness{repo: repo}
}

func (b *findCategoryBusiness) FindCategoryById(
	ctx context.Context,
	id string,
) (*categorymodel.Category, error) {
	data, err := b.repo.FindCategory(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, errors.New("category not found"))
	}

	return data, nil
}
