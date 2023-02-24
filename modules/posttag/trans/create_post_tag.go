package posttagtrans

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	poststore "github.com/iamdevtry/blog/modules/post/store"
	posttagbusiness "github.com/iamdevtry/blog/modules/posttag/business"
	posttagmodel "github.com/iamdevtry/blog/modules/posttag/model"
	posttagrepo "github.com/iamdevtry/blog/modules/posttag/repo"
	posttagstore "github.com/iamdevtry/blog/modules/posttag/store"
	tagstore "github.com/iamdevtry/blog/modules/tag/store"
)

func CreatePostTag(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type reqObj struct {
			PostId string `json:"post_id"`
			TagId  string `json:"tag_id"`
		}
		var req reqObj
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}

		postId, err := common.FromBase58(fmt.Sprintf("%v", req.PostId))
		if err != nil {
			panic(err)
		}

		tagId, err := common.FromBase58(fmt.Sprintf("%v", req.TagId))
		if err != nil {
			panic(err)
		}

		data := posttagmodel.PostTag{
			PostId: int(postId.GetLocalID()),
			TagId:  int(tagId.GetLocalID()),
		}

		store := posttagstore.NewSqlStore(appCtx.GetDBConn())
		tagStore := tagstore.NewSqlStore(appCtx.GetDBConn())
		postStore := poststore.NewSqlStore(appCtx.GetDBConn())
		repo := posttagrepo.NewCreatePostTagRepo(store, tagStore, postStore)
		biz := posttagbusiness.NewCreatePostTagBusiness(repo)
		if err := biz.CreatePostTag(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
