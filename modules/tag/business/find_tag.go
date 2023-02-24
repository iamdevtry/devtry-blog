package tagbusiness

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type FindTagRepo interface {
	FindTag(ctx context.Context, cond map[string]interface{}) (data *tagmodel.Tag, err error)
}

type findTagBusiness struct {
	repo FindTagRepo
}

func NewFindTagBusiness(repo FindTagRepo) *findTagBusiness {
	return &findTagBusiness{repo: repo}
}

func (b *findTagBusiness) FindTagById(ctx context.Context, id int) (data *tagmodel.Tag, err error) {
	data, err = b.repo.FindTag(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(tagmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrCannotGetEntity(tagmodel.EntityName, errors.New("tag not found"))
	}

	return data, nil
}
