package tagtrans

import (
	"github.com/gin-gonic/gin"

	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	tagbusiness "github.com/iamdevtry/blog/modules/tag/business"
	tagrepo "github.com/iamdevtry/blog/modules/tag/repo"
	tagstore "github.com/iamdevtry/blog/modules/tag/store"
)

func ListTag(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := tagstore.NewSqlStore(appCtx.GetDBConn())
		repo := tagrepo.NewListTagRepo(store)
		biz := tagbusiness.NewListTagBusiness(repo)

		result, err := biz.ListTag(c.Request.Context())
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
