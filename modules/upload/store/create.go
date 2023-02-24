package store

import (
	"context"

	"github.com/iamdevtry/blog/common"
)

func (store *sqlStore) Create(ctx context.Context, data *common.Image) error {
	db := store.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
