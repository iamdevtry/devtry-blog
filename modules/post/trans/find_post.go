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
		uuid := c.Param("id")

		store := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := postrepo.NewFindPostRepo(store)
		biz := postbusiness.NewFindPostBusiness(repo)

		data, err := biz.FindPostById(c.Request.Context(), uuid)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func GetPostBySlug(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")

		store := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := postrepo.NewFindPostRepo(store)
		biz := postbusiness.NewFindPostBusiness(repo)

		data, err := biz.FindPostBySlug(c.Request.Context(), slug)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
