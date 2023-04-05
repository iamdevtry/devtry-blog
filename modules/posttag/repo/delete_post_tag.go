package posttagrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
)

type DeletePostTagStore interface {
	DeletePostTag(ctx context.Context, postId, tagId string) error
}

type deletePostTagRepo struct {
	store DeletePostTagStore
}

func NewDeletePostTagRepo(store DeletePostTagStore) *deletePostTagRepo {
	return &deletePostTagRepo{store: store}
}

func (r *deletePostTagRepo) DeletePostTag(ctx context.Context, postId, tagId string) error {
	if err := r.store.DeletePostTag(ctx, postId, tagId); err != nil {
		return common.ErrCannotDeletedEntity(posttagmodel.EntityName, err)
	}

	return nil
}
