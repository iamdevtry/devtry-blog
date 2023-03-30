package tagbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type UpdateTagRepo interface {
	UpdateTag(ctx context.Context, id string, data *tagmodel.TagUpdate) error
}

type updateTagBusiness struct {
	repo UpdateTagRepo
}

func NewUpdateTagBusiness(repo UpdateTagRepo) *updateTagBusiness {
	return &updateTagBusiness{repo: repo}
}

func (b *updateTagBusiness) UpdateTag(ctx context.Context, id string, data *tagmodel.TagUpdate) error {
	if err := b.repo.UpdateTag(ctx, id, data); err != nil {
		return common.ErrCannotUpdatedEntity(tagmodel.EntityName, err)
	}

	return nil
}
