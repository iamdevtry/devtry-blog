package posttrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"

	postbiz "github.com/iamdevtry/blog/modules/post/business"
	postmodel "github.com/iamdevtry/blog/modules/post/model"
	postrepo "github.com/iamdevtry/blog/modules/post/repo"
	poststore "github.com/iamdevtry/blog/modules/post/store"
)

func CreatePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.PostCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.AuthorId = requester.GetUserId()

		if data.FakeParentId != nil {
			uParentId, err := common.FromBase58(data.FakeParentId.String())
			if err != nil {
				panic(err)
			}
			data.ParentId = int(uParentId.GetLocalID())
		}

		store := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := postrepo.NewCreatePostRepo(store)
		biz := postbiz.NewCreatePostBusiness(repo)

		if err := biz.CreatePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
