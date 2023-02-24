package uploadrepo

import (
	"strings"

	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component/uploadprovider"
	uploadmodel "github.com/iamdevtry/blog/modules/upload/model"
	"golang.org/x/net/context"
)

type DeleteStore interface {
	Delete(ctx context.Context, fileName string) error
}

type deleteRepo struct {
	provider    uploadprovider.UploadProvider
	deleteStore DeleteStore
}

func NewDeleteRepo(provider uploadprovider.UploadProvider, deleteStore DeleteStore) *deleteRepo {
	return &deleteRepo{
		provider:    provider,
		deleteStore: deleteStore,
	}
}

func (r *deleteRepo) Delete(ctx context.Context, folder, fileName string) error {
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	key := folder + "/" + fileName

	err := r.provider.DeleteFileUploaded(ctx, key)

	if err != nil {
		return uploadmodel.ErrCannotDeleteFile(err)
	}

	if err := r.deleteStore.Delete(ctx, fileName); err != nil {
		return common.ErrCannotDeletedEntity(uploadmodel.EntityName, err)
	}

	return nil
}
