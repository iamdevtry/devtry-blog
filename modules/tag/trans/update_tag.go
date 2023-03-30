package tagtrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	tagbusiness "github.com/iamdevtry/blog/modules/tag/business"
	tagmodel "github.com/iamdevtry/blog/modules/tag/model"
	tagrepo "github.com/iamdevtry/blog/modules/tag/repo"
	tagstore "github.com/iamdevtry/blog/modules/tag/store"
)

func UpdateTag(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// id, err := strconv.Atoi(c.Param("id"))
		uid := c.Param("id")

		var data tagmodel.TagUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := tagstore.NewSqlStore(appCtx.GetDBConn())
		repo := tagrepo.NewUpdateTagRepo(store)
		biz := tagbusiness.NewUpdateTagBusiness(repo)

		if err := biz.UpdateTag(c.Request.Context(), uid, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
