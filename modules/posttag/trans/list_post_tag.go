package posttagtrans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	posttagbusiness "github.com/iamdevtry/blog/modules/posttag/business"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
	posttagrepo "github.com/iamdevtry/blog/modules/posttag/repo"
	posttagstore "github.com/iamdevtry/blog/modules/posttag/store"
)

func ListPostByTag(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("id")

		filter := posttagmodel.Filter{
			TagId: uuid.MustParse(uid),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := posttagstore.NewSqlStore(appCtx.GetDBConn())
		repo := posttagrepo.NewListPostTagRepo(store)
		biz := posttagbusiness.NewListPostTagBusiness(repo)

		result, err := biz.ListPostTag(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
