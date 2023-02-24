package postcategorybusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
)

type ListPostCategoryRepo interface {
	GetPostCategories(ctx context.Context,
		conditions map[string]interface{},
		filter *postcategorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimplePost, error)
}

type listPostCategoryBusiness struct {
	repo ListPostCategoryRepo
}

func NewListPostCategoryBusiness(repo ListPostCategoryRepo) *listPostCategoryBusiness {
	return &listPostCategoryBusiness{repo: repo}
}

func (b *listPostCategoryBusiness) ListPostCategory(ctx context.Context,
	filter *postcategorymodel.Filter,
	paging *common.Paging,
) ([]common.SimplePost, error) {
	data, err := b.repo.GetPostCategories(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(postcategorymodel.EntityName, err)
	}
	return data, nil
}
