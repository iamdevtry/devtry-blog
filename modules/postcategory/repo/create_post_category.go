package postcategoryrepo

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
)

type CreatePostCategoryStore interface {
	Create(ctx context.Context, data *postcategorymodel.PostCategory) error
}

type FindPostStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error)
}

type FindCategoryStore interface {
	Find(
		ctx context.Context,
		cond map[string]interface{},
	) (*categorymodel.Category, error)
}

type createPostCategoryRepo struct {
	store             CreatePostCategoryStore
	findPostStore     FindPostStore
	findCategoryStore FindCategoryStore
}

func NewCreatePostCategoryRepo(
	store CreatePostCategoryStore,
	findCategoryStore FindCategoryStore,
	findPostStore FindPostStore) *createPostCategoryRepo {
	return &createPostCategoryRepo{
		store:             store,
		findCategoryStore: findCategoryStore,
		findPostStore:     findPostStore,
	}
}

func (r *createPostCategoryRepo) CreatePostCategory(ctx context.Context, data *postcategorymodel.PostCategory) error {
	// Check category exist
	cat, err := r.findCategoryStore.Find(ctx, map[string]interface{}{"id": data.CategoryId})
	if err != nil {
		if err.Error() == common.ErrRecordNotFound.Error() {
			return common.ErrCannotGetEntity(postcategorymodel.EntityName, errors.New("category not found"))
		}

		return err
	}

	if cat.Status == 0 {
		return common.ErrCannotGetEntity(postcategorymodel.EntityName, errors.New("category not found"))
	}

	// Check post exist
	post, err := r.findPostStore.Find(ctx, map[string]interface{}{"id": data.PostId})
	if err != nil {
		if err.Error() == common.ErrRecordNotFound.Error() {
			return common.ErrCannotGetEntity(postcategorymodel.EntityName, errors.New("post not found"))
		}
		return err
	}

	if post.Status == 0 {
		return common.ErrCannotGetEntity(postcategorymodel.EntityName, errors.New("post not found"))
	}

	if err := r.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postcategorymodel.EntityName, err)
	}

	return nil
}
