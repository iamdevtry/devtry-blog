package uploadbusiness

import (
	"context"

	"github.com/iamdevtry/blog/common"
)

type UploadRepo interface {
	Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error)
}

type uploadBusiness struct {
	repo UploadRepo
}

func NewUploadBusiness(repo UploadRepo) *uploadBusiness {
	return &uploadBusiness{repo: repo}
}

func (b *uploadBusiness) UploadImage(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	img, err := b.repo.Upload(ctx, data, folder, fileName)
	if err != nil {
		return nil, err
	}

	return img, nil
}
