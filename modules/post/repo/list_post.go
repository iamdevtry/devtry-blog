package postrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type ListPostStore interface {
	ListPostByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type lisPostRepo struct {
	store ListPostStore
}

func NewListPostRepo(store ListPostStore) *lisPostRepo {
	return &lisPostRepo{store: store}
}

func (r *lisPostRepo) ListPost(
	ctx context.Context,
	paging *common.Paging,
) ([]postmodel.Post, error) {
	result, err := r.store.ListPostByCondition(ctx, nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, nil
}
