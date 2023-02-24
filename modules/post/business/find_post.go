package postbusiness

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

type FindPostRepo interface {
	FindPost(ctx context.Context, cond map[string]interface{}) (data *postmodel.Post, err error)
}

type findPostBusiness struct {
	repo FindPostRepo
}

func NewFindPostBusiness(repo FindPostRepo) *findPostBusiness {
	return &findPostBusiness{repo: repo}
}

func (b *findPostBusiness) FindPostById(ctx context.Context, id int) (data *postmodel.Post, err error) {
	data, err = b.repo.FindPost(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrCannotGetEntity(postmodel.EntityName, errors.New("post not found"))
	}

	return data, nil
}
