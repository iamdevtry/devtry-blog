package postcategorytrans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
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
		repo := postcategoryrepo.NewListPostCategoryRepo(store)
		biz := postcategorybusiness.NewListPostCategoryBusiness(repo)

		result, err := biz.ListPostCategory(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
