package poststore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

func (s *sqlStore) ListPostByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]postmodel.Post, error) {
	var data []postmodel.Post
	db := s.db

	db = db.Table(postmodel.Post{}.TableName()).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if paging.FakeCursor != "" {
		db = db.Where("id <  ?", paging.FakeCursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
