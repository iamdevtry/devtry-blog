package posttagrepo

import (
	"context"
	"errors"

	"github.com/iamdevtry/blog/common"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
)

type CreatePostTagStore interface {
	Create(ctx context.Context, data *posttagmodel.PostTag) error
}

type FindTagStore interface {
	FindTag(ctx context.Context, cond map[string]interface{}) (*tagmodel.Tag, error)
}

type FindPostStore interface {
	Find(ctx context.Context, cond map[string]interface{}) (*postmodel.Post, error)
}

type createPostTagRepo struct {
	store         CreatePostTagStore
	findTagStore  FindTagStore
	findPostStore FindPostStore
}

func NewCreatePostTagRepo(store CreatePostTagStore, findTagStore FindTagStore, findPostStore FindPostStore) *createPostTagRepo {
	return &createPostTagRepo{store: store, findTagStore: findTagStore, findPostStore: findPostStore}
}

func (r *createPostTagRepo) CreatePostTag(ctx context.Context, data *posttagmodel.PostTag) error {
	// check tag exist
	tag, err := r.findTagStore.FindTag(ctx, map[string]interface{}{"id": data.TagId})

	if err != nil {
		if err.Error() == common.ErrRecordNotFound.Error() {
			return common.ErrCannotGetEntity(tagmodel.EntityName, errors.New("tag not found"))
		}
		return err
	}

	if tag.Status == 0 {
		return common.ErrCannotGetEntity(tagmodel.EntityName, errors.New("tag not found"))
	}

	// check post exist
	post, err := r.findPostStore.Find(ctx, map[string]interface{}{"id": data.PostId})

	if err != nil {
		if err.Error() == common.ErrRecordNotFound.Error() {
			return common.ErrCannotGetEntity(postmodel.EntityName, errors.New("post not found"))
		}
		return err
	}

	if post.Status == 0 {
		return common.ErrCannotGetEntity(postmodel.EntityName, errors.New("post not found"))
	}

	// create post tag
	if err := r.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(posttagmodel.EntityName, err)
	}

	return nil
}
