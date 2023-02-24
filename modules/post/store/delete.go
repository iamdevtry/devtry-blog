package poststore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

func (s *sqlStore) SoftDelete(ctx context.Context, id int) error {
	if err := s.db.Table(postmodel.Post{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
