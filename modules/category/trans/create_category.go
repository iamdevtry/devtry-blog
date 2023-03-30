package categorytrans

import (
	"github.com/gin-gonic/gin"

	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	categorybusiness "github.com/iamdevtry/blog/modules/category/business"
	categorymodel "github.com/iamdevtry/blog/modules/category/model"
	categoryrepo "github.com/iamdevtry/blog/modules/category/repo"
	categorystore "github.com/iamdevtry/blog/modules/category/store"
)

func CreateCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.CategoryCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := categorystore.NewSqlStore(appCtx.GetDBConn())
		repo := categoryrepo.NewCreateCategoryRepo(store)
		biz := categorybusiness.NewCreateCategoryBusiness(repo)

		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
