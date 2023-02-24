package postrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type FindPostStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error)
}

type findPostRepo struct {
	store FindPostStore
}

func NewFindPostRepo(store FindPostStore) *findPostRepo {
	return &findPostRepo{store: store}
}

func (r *findPostRepo) FindPost(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error) {
	data, err := r.store.Find(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	return data, nil
}
