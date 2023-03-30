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

func ListPost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := postrepo.NewListPostRepo(store)
		biz := postbusiness.NewListPostBusiness(repo)

		result, err := biz.ListPost(c.Request.Context(), &paging)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
