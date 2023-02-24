package usertrans

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	hasher "github.com/iamdevtry/blog/component/hasher"

	userbusiness "github.com/iamdevtry/blog/modules/user/business"
	usermodel "github.com/iamdevtry/blog/modules/user/model"
	userrepo "github.com/iamdevtry/blog/modules/user/repo"
	userstore "github.com/iamdevtry/blog/modules/user/store"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSqlStore(appCtx.GetDBConn())
		md5 := hasher.NewMd5Hash()
		repo := userrepo.NewRegisterRepo(store, md5)
		biz := userbusiness.NewRegisterBusiness(repo)

		if err := biz.RegisterUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
