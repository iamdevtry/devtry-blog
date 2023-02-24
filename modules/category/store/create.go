package categorystore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

func (s *sqlStore) Create(ctx context.Context, data *categorymodel.CategoryCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
