package posttagtrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/component"
	posttagbusiness "github.com/iamdevtry/blog/modules/posttag/business"
	posttagrepo "github.com/iamdevtry/blog/modules/posttag/repo"
	posttagstore "github.com/iamdevtry/blog/modules/posttag/store"
)

func DeletePostTag(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId := c.Params.ByName("post_id")
		tagId := c.Params.ByName("tag_id")

		store := posttagstore.NewSqlStore(appCtx.GetDBConn())
		repo := posttagrepo.DeletePostTagStore(store)
		biz := posttagbusiness.NewDeletePostTagBusiness(repo)

		if err := biz.DeletePostTag(c.Request.Context(), postId, tagId); err != nil {
			panic(err)
		}
	}
}
