package tagtrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"

	tagbiz "github.com/iamdevtry/blog/modules/tag/business"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
	tagrepo "github.com/iamdevtry/blog/modules/tag/repo"
	tagstore "github.com/iamdevtry/blog/modules/tag/store"
)

func CreateTag(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data tagmodel.TagCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := tagstore.NewSqlStore(appCtx.GetDBConn())
		repo := tagrepo.NewCreateTagRepo(store)
		biz := tagbiz.NewCreateTagBusiness(repo)

		if err := biz.CreateTag(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
