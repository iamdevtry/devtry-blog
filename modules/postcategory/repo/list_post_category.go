package postcategoryrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
)

type ListPostCategoryStore interface {
	ListPostCategoryByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *postcategorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimplePost, error)
}

type listPostCategoryRepo struct {
	store ListPostCategoryStore
}

func NewListPostCategoryRepo(store ListPostCategoryStore) *listPostCategoryRepo {
	return &listPostCategoryRepo{store: store}
}

func (r *listPostCategoryRepo) GetPostCategories(ctx context.Context,
	conditions map[string]interface{},
	filter *postcategorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimplePost, error) {
	data, err := r.store.ListPostCategoryByCondition(ctx, conditions, filter, paging, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(postcategorymodel.EntityName, err)
	}
	return data, nil
}
