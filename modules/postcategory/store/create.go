package postcategorystore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
)

func (s *sqlStore) Create(ctx context.Context, data *postcategorymodel.PostCategory) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
