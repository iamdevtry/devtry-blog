package postbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type CreatePostRepo interface {
	CreatePost(ctx context.Context, data *postmodel.PostCreate) error
}

type createPostBusiness struct {
	repo CreatePostRepo
}

func NewCreatePostBusiness(repo CreatePostRepo) *createPostBusiness {
	return &createPostBusiness{repo: repo}
}

func (b *createPostBusiness) CreatePost(ctx context.Context, data *postmodel.PostCreate) error {
	if err := b.repo.CreatePost(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	return nil
}
