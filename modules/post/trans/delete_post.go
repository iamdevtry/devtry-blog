package posttrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	postbiz "github.com/iamdevtry/blog/modules/post/business"
	postrepo "github.com/iamdevtry/blog/modules/post/repo"
	poststore "github.com/iamdevtry/blog/modules/post/store"
)

func DeletePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := common.FromBase58(c.Param("id"))

		store := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := postrepo.NewDeletePostStore(store)
		biz := postbiz.NewDeletePostBusiness(repo)

		if err := biz.DeletePost(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
