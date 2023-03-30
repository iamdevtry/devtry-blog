package postbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type UpdatePostRepo interface {
	UpdatePost(ctx context.Context, id string, data *postmodel.PostUpdate) error
}

type updatePostBusiness struct {
	repo UpdatePostRepo
}

func NewUpdatePostBusiness(repo UpdatePostRepo) *updatePostBusiness {
	return &updatePostBusiness{repo: repo}
}

func (b *updatePostBusiness) UpdatePost(ctx context.Context, id string, data *postmodel.PostUpdate) error {
	if err := b.repo.UpdatePost(ctx, id, data); err != nil {
		return common.ErrCannotUpdatedEntity(postmodel.EntityName, err)
	}

	return nil
}
