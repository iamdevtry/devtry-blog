package tagrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type FindTagStore interface {
	FindTag(ctx context.Context, cond map[string]interface{}) (*tagmodel.Tag, error)
}

type findTagRepo struct {
	store FindTagStore
}

func NewFindTagRepo(store FindTagStore) *findTagRepo {
	return &findTagRepo{store: store}
}

func (r *findTagRepo) FindTag(ctx context.Context, cond map[string]interface{}) (*tagmodel.Tag, error) {
	data, err := r.store.FindTag(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotGetEntity(tagmodel.EntityName, err)
	}

	return data, nil
}
