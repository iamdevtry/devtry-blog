package tagbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type DeleteTagRepo interface {
	SoftDeleteTag(ctx context.Context, id string) error
}

type deleteTagBusiness struct {
	repo DeleteTagRepo
}

func NewDeleteTagBusiness(repo DeleteTagRepo) *deleteTagBusiness {
	return &deleteTagBusiness{repo: repo}
}

func (b *deleteTagBusiness) DeleteTag(ctx context.Context, id string) error {
	if err := b.repo.SoftDeleteTag(ctx, id); err != nil {
		return common.ErrCannotDeletedEntity(tagmodel.EntityName, err)
	}

	return nil
}
