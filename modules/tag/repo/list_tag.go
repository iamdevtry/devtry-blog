package tagrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type ListTagStore interface {
	ListTagByCondition(
		ctx context.Context,
	) ([]tagmodel.Tag, error)
}

type lisTagRepo struct {
	store ListTagStore
}

func NewListTagRepo(store ListTagStore) *lisTagRepo {
	return &lisTagRepo{store: store}
}

func (r *lisTagRepo) ListTag(ctx context.Context) ([]tagmodel.Tag, error) {
	result, err := r.store.ListTagByCondition(ctx)

	if err != nil {
		return nil, common.ErrCannotListEntity(tagmodel.EntityName, err)
	}

	return result, nil
}
