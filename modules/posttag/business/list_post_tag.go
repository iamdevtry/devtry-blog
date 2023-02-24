package posttagbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
)

type ListPostCategoryRepo interface {
	GetPostTag(ctx context.Context,
		conditions map[string]interface{},
		filter *posttagmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimplePost, error)
}

type listPostTagBusiness struct {
	repo ListPostCategoryRepo
}

func NewListPostTagBusiness(repo ListPostCategoryRepo) *listPostTagBusiness {
	return &listPostTagBusiness{repo: repo}
}

func (b *listPostTagBusiness) ListPostTag(ctx context.Context,
	filter *posttagmodel.Filter,
	paging *common.Paging,
) ([]common.SimplePost, error) {
	data, err := b.repo.GetPostTag(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(posttagmodel.EntityName, err)
	}
	return data, nil
}
