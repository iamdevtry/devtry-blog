package tagstore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

func (s *sqlStore) SoftDelete(ctx context.Context, id string) error {
	if err := s.db.Table(tagmodel.Tag{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
