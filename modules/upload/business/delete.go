package uploadbusiness

import (
	"context"
)

type DeleteRepo interface {
	Delete(ctx context.Context, folder, fileName string) error
}

type deleteBusiness struct {
	repo DeleteRepo
}

func NewDeleteBusiness(repo DeleteRepo) *deleteBusiness {
	return &deleteBusiness{repo: repo}
}

func (b *deleteBusiness) DeleteImage(ctx context.Context, folder, fileName string) error {
	err := b.repo.Delete(ctx, folder, fileName)
	if err != nil {
		return err
	}

	return nil
}
