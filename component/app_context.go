package component

import (
	"github.com/iamdevtry/blog/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetDBConn() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appContext struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) *appContext {
	return &appContext{
		db:         db,
		upProvider: upProvider,
		secretKey:  secretKey,
	}
}

func (appCtx *appContext) GetDBConn() *gorm.DB {
	return appCtx.db
}
func (appCtx *appContext) SecretKey() string {
	return appCtx.secretKey
}
func (appCtx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appCtx.upProvider
}
