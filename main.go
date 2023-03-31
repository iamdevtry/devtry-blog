package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iamdevtry/blog/component"
	"github.com/iamdevtry/blog/component/uploadprovider"
	"github.com/iamdevtry/blog/middleware"
	"github.com/iamdevtry/blog/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	categorytrans "github.com/iamdevtry/blog/modules/category/trans"
	posttrans "github.com/iamdevtry/blog/modules/post/trans"
	postcategorytrans "github.com/iamdevtry/blog/modules/postcategory/trans"
	posttagtrans "github.com/iamdevtry/blog/modules/posttag/trans"
	tagtrans "github.com/iamdevtry/blog/modules/tag/trans"
	uploadtrans "github.com/iamdevtry/blog/modules/upload/trans"
	usertrans "github.com/iamdevtry/blog/modules/user/trans"
)

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {
	appContext := component.NewAppContext(db, upProvider, secretKey)

	route := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	route.Use(cors.New(config), middleware.Recover(appContext))

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := route.Group("/v1")

	v1.POST("/upload", uploadtrans.Upload(appContext))
	v1.DELETE("/delete/:id", uploadtrans.Delete(appContext))

	v1.POST("/register", usertrans.Register(appContext))
	v1.POST("/login", usertrans.Login(appContext))
	v1.GET("/profile", middleware.RequireAuth(appContext), usertrans.GetProfile(appContext))

	categories := v1.Group("/categories")
	{
		categories.GET("", categorytrans.ListCategory(appContext))
		categories.POST("", middleware.RequireAuth(appContext), categorytrans.CreateCategory(appContext))
		categories.GET("/:id", categorytrans.FindCategoryById(appContext))
		categories.PUT("/:id", middleware.RequireAuth(appContext), categorytrans.UpdateCategoryById(appContext))
		categories.DELETE("/:id", middleware.RequireAuth(appContext), categorytrans.DeleteCategory(appContext))

		categories.POST("/add-post", middleware.RequireAuth(appContext), postcategorytrans.CreatePostCategory(appContext))
		categories.GET("/:id/posts", postcategorytrans.ListPostByCategory(appContext))
	}

	posts := v1.Group("/posts")
	{
		posts.GET("", posttrans.ListPost(appContext))
		posts.POST("", middleware.RequireAuth(appContext), posttrans.CreatePost(appContext))
		posts.GET("/:id", posttrans.GetPostById(appContext))
		posts.PUT("/:id", middleware.RequireAuth(appContext), posttrans.UpdatePost(appContext))
		posts.DELETE("/:id", middleware.RequireAuth(appContext), posttrans.DeletePost(appContext))
	}

	tags := v1.Group("/tags")
	{
		tags.GET("", tagtrans.ListTag(appContext))
		tags.POST("", middleware.RequireAuth(appContext), tagtrans.CreateTag(appContext))
		tags.GET("/:id", tagtrans.GetTagById(appContext))
		tags.PUT("/:id", middleware.RequireAuth(appContext), tagtrans.UpdateTag(appContext))
		tags.DELETE("/:id", middleware.RequireAuth(appContext), tagtrans.DeleteTag(appContext))

		tags.POST("/add-post", middleware.RequireAuth(appContext), posttagtrans.CreatePostTag(appContext))
		tags.GET("/:id/posts", posttagtrans.ListPostByCategory(appContext))
	}

	return route.Run()
}

func main() {
	//load config
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config environment", err)
	}

	s3Provider := uploadprovider.NewS3Provider(config.S3BucketName, config.S3Region, config.S3APIKey, config.S3Secret, config.S3Domain)

	//connect to database
	sqlDB, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	if err := runService(db, s3Provider, config.SysSecretKey); err != nil {
		log.Fatal("cannot run service:", err)
	}
}
