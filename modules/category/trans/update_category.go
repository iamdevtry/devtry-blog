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

func UpdateCategoryById(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := c.Param("id")

		var data categorymodel.CategoryUpdate
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		store := categorystore.NewSqlStore(appCtx.GetDBConn())
		repo := categoryrepo.NewUpdateCategoryRepo(store)
		biz := categorybusiness.NewUpdateCategoryBusiness(repo)

		if err := biz.UpdateCategory(c.Request.Context(), uuid, &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
