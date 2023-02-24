package postbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type ListPostRepo interface {
	ListPost(
		ctx context.Context,
		paging *common.Paging,
	) ([]postmodel.Post, error)
}

type listPostBusiness struct {
	repo ListPostRepo
}

func NewListPostBusiness(repo ListPostRepo) *listPostBusiness {
	return &listPostBusiness{repo: repo}
}

func (b *listPostBusiness) ListPost(
	ctx context.Context,
	paging *common.Paging,
) ([]postmodel.Post, error) {

	result, err := b.repo.ListPost(ctx, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, nil
}
