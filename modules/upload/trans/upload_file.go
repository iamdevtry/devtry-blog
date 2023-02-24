package uploadtrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"

	_ "image/jpeg"
	_ "image/png"

	uploadbusiness "github.com/iamdevtry/blog/modules/upload/business"
	uploadrepo "github.com/iamdevtry/blog/modules/upload/repo"
	uploadstore "github.com/iamdevtry/blog/modules/upload/store"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstore.NewSqlStore(appCtx.GetDBConn())
		repo := uploadrepo.NewUploadRepo(appCtx.UploadProvider(), imgStore)
		biz := uploadbusiness.NewUploadBusiness(repo)
		img, err := biz.UploadImage(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(img))
	}
}
