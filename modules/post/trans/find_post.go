package posttrans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	postbusiness "github.com/iamdevtry/blog/modules/post/business"
	postrepo "github.com/iamdevtry/blog/modules/post/repo"
	poststore "github.com/iamdevtry/blog/modules/post/store"
)

func GetPostById(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := common.FromBase58(c.Param("id"))

		store := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := postrepo.NewFindPostRepo(store)
		biz := postbusiness.NewFindPostBusiness(repo)

		data, err := biz.FindPostById(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
