package tagbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type ListTagRepo interface {
	ListTag(ctx context.Context) ([]tagmodel.Tag, error)
}

type listTagBusiness struct {
	repo ListTagRepo
}

func NewListTagBusiness(repo ListTagRepo) *listTagBusiness {
	return &listTagBusiness{repo: repo}
}

func (b *listTagBusiness) ListTag(ctx context.Context) ([]tagmodel.Tag, error) {

	result, err := b.repo.ListTag(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(tagmodel.EntityName, err)
	}

	return result, nil
}
