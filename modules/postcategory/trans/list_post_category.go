package postcategorytrans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	postcategorybusiness "github.com/iamdevtry/blog/modules/postcategory/business"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
	postcategoryrepo "github.com/iamdevtry/blog/modules/postcategory/repo"
	postcategorystore "github.com/iamdevtry/blog/modules/postcategory/store"
)

func ListPostByCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		// id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := postcategorymodel.Filter{
			CategoryId: int(uid.GetLocalID()),
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

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
