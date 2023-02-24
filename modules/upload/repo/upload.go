package uploadrepo

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component/uploadprovider"
	"golang.org/x/net/context"

	uploadmodel "github.com/iamdevtry/blog/modules/upload/model"
)

type CreateImageStore interface {
	Create(ctx context.Context, data *common.Image) error
}

type uploadRepo struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStore
}

func NewUploadRepo(provider uploadprovider.UploadProvider, imgStore CreateImageStore) *uploadRepo {
	return &uploadRepo{
		provider: provider,
		imgStore: imgStore,
	}
}

func (r *uploadRepo) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := r.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.CloudName = "s3"
	img.Extension = fileExt
	img.FileName = fileName

	if err := r.imgStore.Create(ctx, img); err != nil {
		return nil, common.ErrCannotCreateEntity(uploadmodel.EntityName, err)
	}

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
