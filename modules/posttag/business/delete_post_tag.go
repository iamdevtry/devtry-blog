package posttagbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
)

type DeletePostTagRepo interface {
	DeletePostTag(ctx context.Context, postId, tagId string) error
}

type deletePostTagBusiness struct {
	repo DeletePostTagRepo
}

func NewDeletePostTagBusiness(repo DeletePostTagRepo) *deletePostTagBusiness {
	return &deletePostTagBusiness{repo: repo}
}

func (b *deletePostTagBusiness) DeletePostTag(ctx context.Context, postId, tagId string) error {
	if err := b.repo.DeletePostTag(ctx, postId, tagId); err != nil {
		return common.ErrCannotDeletedEntity(posttagmodel.EntityName, err)
	}

	return nil
}
