package tagrepo

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type UpdateTagStore interface {
	FindTag(ctx context.Context, cond map[string]interface{}) (*tagmodel.Tag, error)
	Update(ctx context.Context, id int, data *tagmodel.TagUpdate) error
}

type updateTagRepo struct {
	store UpdateTagStore
}

func NewUpdateTagRepo(store UpdateTagStore) *updateTagRepo {
	return &updateTagRepo{store: store}
}

func (r *updateTagRepo) UpdateTag(ctx context.Context, id int, data *tagmodel.TagUpdate) error {
	oldData, err := r.store.FindTag(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrCannotGetEntity(tagmodel.EntityName, errors.New("tag not found"))
	}

	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(tagmodel.EntityName, err)
	}

	if err := r.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdatedEntity(tagmodel.EntityName, err)
	}

	return nil
}
