package categorystore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

func (s *sqlStore) Update(ctx context.Context, id string, data *categorymodel.CategoryUpdate) error {
	if err := s.db.Table(categorymodel.CategoryUpdate{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"parent_id":  data.ParentId,
			"meta_title": data.MetaTitle,
			"content":    data.Content,
		}).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
