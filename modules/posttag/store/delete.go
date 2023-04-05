package posttagstore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
)

func (s *sqlStore) DeletePostTag(ctx context.Context, postId, tagId string) error {
	if err := s.db.Table(posttagmodel.PostTag{}.TableName()).
		Where("tag_id = ? and post_id = ?", tagId, postId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
