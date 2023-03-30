package tagtrans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	tagbusiness "github.com/iamdevtry/blog/modules/tag/business"
	tagrepo "github.com/iamdevtry/blog/modules/tag/repo"
	tagstore "github.com/iamdevtry/blog/modules/tag/store"
)

func GetTagById(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// id, err := strconv.Atoi(c.Param("id"))
		uid := c.Param("id")

		store := tagstore.NewSqlStore(appCtx.GetDBConn())
		repo := tagrepo.NewFindTagRepo(store)
		biz := tagbusiness.NewFindTagBusiness(repo)

		data, err := biz.FindTagById(c.Request.Context(), uid)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
