package usertrans

import (
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/common"
	"github.com/iamdevtry/blog/component"
	hasher "github.com/iamdevtry/blog/component/hasher"
	"github.com/iamdevtry/blog/component/tokenprovider/jwt"
	userbusiness "github.com/iamdevtry/blog/modules/user/business"
	usermodel "github.com/iamdevtry/blog/modules/user/model"
	userrepo "github.com/iamdevtry/blog/modules/user/repo"
	userstore "github.com/iamdevtry/blog/modules/user/store"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSqlStore(appCtx.GetDBConn())
		md5 := hasher.NewMd5Hash()
		repo := userrepo.NewLoginRepo(store, tokenProvider, md5, 60*60*24*7)
		business := userbusiness.NewLoginBusiness(repo)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(account))
	}
}
