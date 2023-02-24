package uploadprovider

import (
	"context"

	"github.com/iamdevtry/blog/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
	DeleteFileUploaded(ctx context.Context, key string) error
}
