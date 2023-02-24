package posttagbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
)

type CreatePostTagRepo interface {
	CreatePostTag(ctx context.Context, data *posttagmodel.PostTag) error
}

type createPostTagBusiness struct {
	repo CreatePostTagRepo
}

func NewCreatePostTagBusiness(repo CreatePostTagRepo) *createPostTagBusiness {
	return &createPostTagBusiness{repo: repo}
}

func (b *createPostTagBusiness) CreatePostTag(ctx context.Context, data *posttagmodel.PostTag) error {
	if err := b.repo.CreatePostTag(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(posttagmodel.EntityName, err)
	}
	return nil
}
