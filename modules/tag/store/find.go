package tagstore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

func (s *sqlStore) FindTag(ctx context.Context, cond map[string]interface{}) (*tagmodel.Tag, error) {
	var data tagmodel.Tag

	if err := s.db.Table(tagmodel.TagCreate{}.TableName()).
		Where(cond).
		First(&data).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
