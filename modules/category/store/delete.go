package categorystore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

func (s *sqlStore) SoftDelete(ctx context.Context, id string) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
