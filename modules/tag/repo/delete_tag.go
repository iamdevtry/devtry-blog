package tagrepo

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type DeleteTagStore interface {
	FindTag(ctx context.Context, cond map[string]interface{}) (*tagmodel.Tag, error)
	SoftDelete(ctx context.Context, id int) error
}

type deleteTagRepo struct {
	store DeleteTagStore
}

func NewDeleteTagStore(store DeleteTagStore) *deleteTagRepo {
	return &deleteTagRepo{store: store}
}

func (repo *deleteTagRepo) SoftDeleteTag(ctx context.Context, id int) error {
	oldData, err := repo.store.FindTag(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(tagmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrCannotGetEntity(tagmodel.EntityName, errors.New("tag not found"))
	}

	if err := repo.store.SoftDelete(ctx, id); err != nil {
		return common.ErrCannotDeletedEntity(tagmodel.EntityName, err)
	}

	return nil
}
