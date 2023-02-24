package tagtrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	tagbiz "github.com/iamdevtry/blog/modules/tag/business"
	tagrepo "github.com/iamdevtry/blog/modules/tag/repo"
	tagstore "github.com/iamdevtry/blog/modules/tag/store"
)

func DeleteTag(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// id, err := strconv.Atoi(c.Param("id"))
		uid, _ := common.FromBase58(c.Param("id"))

		store := tagstore.NewSqlStore(appCtx.GetDBConn())
		repo := tagrepo.NewDeleteTagStore(store)
		biz := tagbiz.NewDeleteTagBusiness(repo)

		if err := biz.DeleteTag(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
