package postrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type UpdatePostStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error)
	Update(ctx context.Context, id string, data *postmodel.PostUpdate) error
}

type updatePostRepo struct {
	store UpdatePostStore
}

func NewUpdatePostRepo(store UpdatePostStore) *updatePostRepo {
	return &updatePostRepo{store: store}
}

func (r *updatePostRepo) UpdatePost(ctx context.Context, id string, data *postmodel.PostUpdate) error {
	oldData, err := r.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrCannotGetEntity(postmodel.EntityName, errors.New("post not found"))
	}

	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	if data.ParentId != nil && data.ParentId != oldData.ParentId {
		parentPost, err := r.store.Find(ctx, map[string]interface{}{"id": data.ParentId})

		if err != nil {
			return common.ErrCannotGetEntity(postmodel.EntityName, fmt.Errorf("cannot find post category with id %d", data.ParentId))
		}

		if parentPost.Status == 0 {
			return common.ErrCannotGetEntity(postmodel.EntityName, errors.New("post category not found"))
		}
	}

	if err := r.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdatedEntity(postmodel.EntityName, err)
	}

	return nil
}
