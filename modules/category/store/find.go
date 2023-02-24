package categorystore

import (
	"context"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
)

func (s *sqlStore) Find(
	ctx context.Context,
	cond map[string]interface{},
) (*categorymodel.Category, error) {
	var data categorymodel.Category

	if err := s.db.Table(categorymodel.Category{}.TableName()).
		Where(cond).
		First(&data).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
