package categorytrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	categorybusiness "github.com/iamdevtry/blog/modules/category/business"
	categoryrepo "github.com/iamdevtry/blog/modules/category/repo"
	categorystore "github.com/iamdevtry/blog/modules/category/store"
)

func FindCategoryById(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")

		store := categorystore.NewSqlStore(appCtx.GetDBConn())
		repo := categoryrepo.NewFindCategoryRepo(store)
		biz := categorybusiness.NewFindCategoryBusiness(repo)

		result, err := biz.FindCategoryById(c.Request.Context(), uuid)
		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
