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
	store        FindPostStore
	listTagStore ListTagByPostIdStore
}

func NewFindPostRepo(store FindPostStore, listTagStore ListTagByPostIdStore) *findPostRepo {
	return &findPostRepo{store: store, listTagStore: listTagStore}
}

func (r *findPostRepo) FindPost(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error) {
	data, err := r.store.Find(ctx, cond)

	if err != nil {
		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	tags, _ := r.listTagStore.GetTagsByPostId(ctx, data.Id.String())
	data.Tags = tags

	return data, nil
}
