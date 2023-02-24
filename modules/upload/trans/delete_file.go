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

func Delete(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileName := c.Param("id")
		// folder := c.Param("folder")
		folder := "img" //hardcode

		store := uploadstore.NewSqlStore(appCtx.GetDBConn())
		repo := uploadrepo.NewDeleteRepo(appCtx.UploadProvider(), store)
		biz := uploadbusiness.NewDeleteBusiness(repo)
		err := biz.DeleteImage(c.Request.Context(), folder, fileName)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse("Delete file success"))
	}
}
