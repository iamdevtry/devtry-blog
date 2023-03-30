package postrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type CreatePostStore interface {
	Create(ctx context.Context, data *postmodel.PostCreate) error
	Find(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error)
}

type createPostRepo struct {
	store CreatePostStore
}

func NewCreatePostRepo(store CreatePostStore) *createPostRepo {
	return &createPostRepo{store: store}
}

func (r *createPostRepo) CreatePost(ctx context.Context, data *postmodel.PostCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	if data.ParentId != nil {
		parentPost, err := r.store.Find(ctx, map[string]interface{}{"id": data.ParentId})

		if err != nil {
			return common.ErrCannotGetEntity(postmodel.EntityName, fmt.Errorf("cannot find parent post with id %d", data.ParentId))
		}

		if parentPost.Status == 0 {
			return common.ErrCannotGetEntity(postmodel.EntityName, errors.New("parent post not found"))
		}
	}

	if err := r.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	return nil
}
