package postcategoryrepo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
	categoryrepo "github.com/iamdevtry/blog/modules/category/repo"
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
	store        ListPostCategoryStore
	findCatStore FindCategoryStore
}

func NewListPostCategoryRepo(store ListPostCategoryStore, findCatStore FindCategoryStore) *listPostCategoryRepo {
	return &listPostCategoryRepo{
		store:        store,
		findCatStore: categoryrepo.FindCategoryStore(findCatStore),
	}
}

func (r *listPostCategoryRepo) GetPostCategories(ctx context.Context,
	conditions map[string]interface{},
	filter *postcategorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimplePost, error) {
	var catId uuid.UUID
	if conditions != nil {
		if slug, ok := conditions["slug"]; ok {
			cat, err := r.findCatStore.Find(ctx, map[string]interface{}{"slug": slug})

			if err != nil {
				return nil, common.ErrCannotListEntity(postcategorymodel.EntityName, fmt.Errorf("cannot find category with slug %s", slug))
			} else {
				catId = cat.Id
			}
		}
	}

	if catId != uuid.Nil {
		filter = &postcategorymodel.Filter{
			CategoryId: catId,
		}
	}

	data, err := r.store.ListPostCategoryByCondition(ctx, nil, filter, paging, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(postcategorymodel.EntityName, err)
	}
	return data, nil
}
