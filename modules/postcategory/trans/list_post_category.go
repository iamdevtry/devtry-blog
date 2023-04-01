package postcategorytrans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	categorystore "github.com/iamdevtry/blog/modules/category/store"
	postcategorybusiness "github.com/iamdevtry/blog/modules/postcategory/business"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
	postcategoryrepo "github.com/iamdevtry/blog/modules/postcategory/repo"
	postcategorystore "github.com/iamdevtry/blog/modules/postcategory/store"
)

func ListPostByCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("id")
		// id, err := strconv.Atoi(c.Param("id"))

		filter := postcategorymodel.Filter{
			CategoryId: uuid.MustParse(uid),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := postcategorystore.NewSqlStore(appCtx.GetDBConn())
		repo := postcategoryrepo.NewListPostCategoryRepo(store, nil)
		biz := postcategorybusiness.NewListPostCategoryBusiness(repo)

		result, err := biz.ListPostCategory(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}

func ListPostByCategorySlug(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		// id, err := strconv.Atoi(c.Param("id"))

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := postcategorystore.NewSqlStore(appCtx.GetDBConn())
		catStore := categorystore.NewSqlStore(appCtx.GetDBConn())
		repo := postcategoryrepo.NewListPostCategoryRepo(store, catStore)
		biz := postcategorybusiness.NewListPostCategoryBusiness(repo)

		result, err := biz.ListPostCategoryBySlug(c.Request.Context(), slug, nil, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
