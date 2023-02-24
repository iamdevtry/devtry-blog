package postcategorytrans

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	categorystore "github.com/iamdevtry/blog/modules/category/store"
	poststore "github.com/iamdevtry/blog/modules/post/store"
	postcategorybusiness "github.com/iamdevtry/blog/modules/postcategory/business"
	postcategorymodel "github.com/iamdevtry/blog/modules/postcategory/model"
	postcategoryrepo "github.com/iamdevtry/blog/modules/postcategory/repo"
	postcategorystore "github.com/iamdevtry/blog/modules/postcategory/store"
)

func CreatePostCategory(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type reqObj struct {
			PostId     string `json:"post_id"`
			CategoryId string `json:"category_id"`
		}
		var req reqObj
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}

		postId, err := common.FromBase58(fmt.Sprintf("%v", req.PostId))
		if err != nil {
			panic(err)
		}

		catId, err := common.FromBase58(fmt.Sprintf("%v", req.CategoryId))
		if err != nil {
			panic(err)
		}

		data := postcategorymodel.PostCategory{
			PostId:     int(postId.GetLocalID()),
			CategoryId: int(catId.GetLocalID()),
		}

		store := postcategorystore.NewSqlStore(appCtx.GetDBConn())
		postStore := poststore.NewSqlStore(appCtx.GetDBConn())
		catStore := categorystore.NewSqlStore(appCtx.GetDBConn())
		repo := postcategoryrepo.NewCreatePostCategoryRepo(store, catStore, postStore)
		biz := postcategorybusiness.NewCreatePostCategoryBusiness(repo)
		if err := biz.CreatePostCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
