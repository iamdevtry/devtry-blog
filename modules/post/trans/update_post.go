package posttrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	postbusiness "github.com/iamdevtry/blog/modules/post/business"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
	postrepo "github.com/iamdevtry/blog/modules/post/repo"
	poststore "github.com/iamdevtry/blog/modules/post/store"
)

func UpdatePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")

		var data postmodel.PostUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := postrepo.NewUpdatePostRepo(store)
		biz := postbusiness.NewUpdatePostBusiness(repo)

		if err := biz.UpdatePost(c.Request.Context(), uuid, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
