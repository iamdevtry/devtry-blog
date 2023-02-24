package poststore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

func (s *sqlStore) Find(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error) {
	var data postmodel.Post

	if err := s.db.Table(postmodel.Post{}.TableName()).
		Where(cond).
		First(&data).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
