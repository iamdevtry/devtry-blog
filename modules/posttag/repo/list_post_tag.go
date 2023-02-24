package posttagrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
)

type ListPostTagStore interface {
	ListPostTagByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *posttagmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimplePost, error)
}

type ListPostTagRepo struct {
	store ListPostTagStore
}

func NewListPostTagRepo(store ListPostTagStore) *ListPostTagRepo {
	return &ListPostTagRepo{store: store}
}

func (r *ListPostTagRepo) GetPostTag(ctx context.Context,
	conditions map[string]interface{},
	filter *posttagmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimplePost, error) {
	data, err := r.store.ListPostTagByCondition(ctx, conditions, filter, paging, moreKeys...)
	if err != nil {
		return nil, common.ErrCannotListEntity(posttagmodel.EntityName, err)
	}
	return data, nil
}
