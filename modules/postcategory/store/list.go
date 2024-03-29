package postcategorystore

import (
	"context"

	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
)

// const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) ListPostCategoryByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *postcategorymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimplePost, error) {
	var result []postcategorymodel.PostCategory
	db := s.db
	db.Debug()

	db = db.Table(postcategorymodel.PostCategory{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.CategoryId != uuid.Nil {
			db = db.Where("category_id = ?", v.CategoryId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Preload("Post")

	// if v := paging.FakeCursor; v != "" {
	// 	timeCreated, err := time.Parse(timeLayout, string(v))

	// 	if err != nil {
	// 		return nil, common.ErrDB(err)
	// 	}

	// 	db = db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	// } else {
	// 	db = db.Offset((paging.Page - 1) * paging.Limit)
	// }

	if err := db.
		Limit(paging.Limit).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	posts := make([]common.SimplePost, len(result))

	for i := range result {
		// result[i].User.CreatedAt = item.CreatedAt
		// result[i].User.UpdatedAt = nil
		posts[i] = *result[i].Post

		// if i == len(result)-1 {
		// 	cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
		// 	paging.NextCursor = cursorStr
		// }
	}

	return posts, nil
}
