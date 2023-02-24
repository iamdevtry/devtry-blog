package tagrepo

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type CreateTagStore interface {
	Create(ctx context.Context, data *tagmodel.TagCreate) error
}

type createTagRepo struct {
	store CreateTagStore
}

func NewCreateTagRepo(store CreateTagStore) *createTagRepo {
	return &createTagRepo{store: store}
}

func (r *createTagRepo) CreateTag(ctx context.Context, data *tagmodel.TagCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(tagmodel.EntityName, err)
	}

	if err := r.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(tagmodel.EntityName, err)
	}

	return nil
}
