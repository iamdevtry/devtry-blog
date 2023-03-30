package postrepo

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type DeletePostStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error)
	SoftDelete(ctx context.Context, id string) error
}

type deletePostRepo struct {
	store DeletePostStore
}

func NewDeletePostStore(store DeletePostStore) *deletePostRepo {
	return &deletePostRepo{store: store}
}

func (repo *deletePostRepo) SoftDeletePost(ctx context.Context, id string) error {
	oldData, err := repo.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrCannotGetEntity(postmodel.EntityName, errors.New("post not found"))
	}

	if err := repo.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeletedEntity(postmodel.EntityName, err)
	}

	return nil
}
