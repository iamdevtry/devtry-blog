package tagbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type CreateTagRepo interface {
	CreateTag(ctx context.Context, data *tagmodel.TagCreate) error
}

type createTagBusiness struct {
	repo CreateTagRepo
}

func NewCreateTagBusiness(repo CreateTagRepo) *createTagBusiness {
	return &createTagBusiness{repo: repo}
}

func (b *createTagBusiness) CreateTag(ctx context.Context, data *tagmodel.TagCreate) error {
	if err := b.repo.CreateTag(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(tagmodel.EntityName, err)
	}

	return nil
}
