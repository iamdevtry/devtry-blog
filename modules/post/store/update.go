package poststore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
)

func (s *sqlStore) Update(ctx context.Context, id string, data *postmodel.PostUpdate) error {
	if err := s.db.Table(postmodel.PostUpdate{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"parent_id":  data.ParentId,
			"meta_title": data.MetaTitle,
			"sumary":     data.Sumary,
			"content":    data.Content,
			"thumbnail":  data.Thumbnail,
		}).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
