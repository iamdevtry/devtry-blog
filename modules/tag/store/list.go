package tagstore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

func (s *sqlStore) ListTagByCondition(
	ctx context.Context,
) ([]tagmodel.Tag, error) {
	var data []tagmodel.Tag

	if err := s.db.Table(tagmodel.Tag{}.TableName()).Where("status in (1)").Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
