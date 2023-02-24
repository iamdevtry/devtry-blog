package store

import (
	"context"

	"github.com/iamdevtry/blog/common"
	uploadmodel "github.com/iamdevtry/blog/modules/upload/model"
)

func (sql *sqlStore) Delete(ctx context.Context, fileName string) error {
	if err := sql.db.Table(uploadmodel.Upload{}.TableName()).Where("file_name = ?", fileName).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
