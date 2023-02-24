package postbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type DeletePostRepo interface {
	SoftDeletePost(ctx context.Context, id int) error
}

type deletePostBusiness struct {
	repo DeletePostRepo
}

func NewDeletePostBusiness(repo DeletePostRepo) *deletePostBusiness {
	return &deletePostBusiness{repo: repo}
}

func (b *deletePostBusiness) DeletePost(ctx context.Context, id int) error {
	if err := b.repo.SoftDeletePost(ctx, id); err != nil {
		return common.ErrCannotDeletedEntity(postmodel.EntityName, err)
	}

	return nil
}
