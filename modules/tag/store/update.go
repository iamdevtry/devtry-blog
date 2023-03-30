package tagstore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

func (s *sqlStore) Update(ctx context.Context, id string, data *tagmodel.TagUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
