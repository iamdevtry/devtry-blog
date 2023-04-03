package postrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type ListPostStore interface {
	ListPostByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type ListTagByPostIdStore interface {
	GetTagsByPostId(ctx context.Context, id string) ([]tagmodel.Tag, error)
}

type lisPostRepo struct {
	store        ListPostStore
	listTagStore ListTagByPostIdStore
}

func NewListPostRepo(store ListPostStore, listTagStore ListTagByPostIdStore) *lisPostRepo {
	return &lisPostRepo{store: store, listTagStore: listTagStore}
}

func (r *lisPostRepo) ListPost(
	ctx context.Context,
	paging *common.Paging,
) ([]postmodel.Post, error) {
	result, err := r.store.ListPostByCondition(ctx, nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	for i := 0; i < len(result); i++ {
		tags, _ := r.listTagStore.GetTagsByPostId(ctx, result[i].Id.String())
		result[i].Tags = tags
	}

	return result, nil
}
