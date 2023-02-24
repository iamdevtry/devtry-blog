package categorytrans

import (
	"github.com/gin-gonic/gin"

	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	categorybusiness "github.com/iamdevtry/blog/modules/category/business"
	categoryrepo "github.com/iamdevtry/blog/modules/category/repo"
	categorystore "github.com/iamdevtry/blog/modules/category/store"
)

func ListCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := categorystore.NewSqlStore(appCtx.GetDBConn())
		repo := categoryrepo.NewListCategoryRepo(store)
		biz := categorybusiness.NewListCategoryBusiness(repo)

		result, err := biz.ListCategory(c.Request.Context())
		if err != nil {
			panic(err)
		}

		//Mask id to uid
		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
