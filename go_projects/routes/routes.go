package routes

import (
	"blog-system/config"
	"blog-system/controllers"
	"blog-system/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 中间件
	r.Use(middleware.CORSMiddleware())

	// 控制器
	authController := controllers.NewAuthController()
	articleController := controllers.NewArticleController()
	musicController := controllers.NewMusicController()
	categoryController := controllers.NewCategoryController()
	tagController := controllers.NewTagController()
	commentController := controllers.NewCommentController()
	linkController := controllers.NewLinkController()
	uploadController := controllers.NewUploadController()
	labController := controllers.NewLabController()

	// 公开路由
	api := r.Group("/api")
	{
		// 认证
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

		// 文章
		articles := api.Group("/articles")
		{
			articles.GET("", articleController.GetArticles)
			articles.GET("/:id", articleController.GetArticle)
			articles.POST("/:id/like", articleController.LikeArticle)
		}

		// 分类
		categories := api.Group("/categories")
		{
			categories.GET("", categoryController.GetCategories)
			categories.GET("/:id", categoryController.GetCategory)
		}

		// 标签
		tags := api.Group("/tags")
		{
			tags.GET("", tagController.GetTags)
		}

		// 评论
		comments := api.Group("/comments")
		{
			comments.GET("", commentController.GetComments)
			comments.POST("", commentController.CreateComment)
		}

			// 音乐
		music := api.Group("/music")
		{
			music.GET("", musicController.GetMusics)
			music.GET("/:id", musicController.GetMusic)
			music.GET("/playlists", musicController.GetPlaylists)
			music.GET("/playlists/:id", musicController.GetPlaylist)
		}

		// 友情链接
		links := api.Group("/links")
		{
			links.GET("", linkController.GetLinks)
		}

		// 实验室模块
		labs := api.Group("/labs")
		{
			labs.GET("", labController.GetLabs)
			labs.GET("/:slug", labController.GetLab)
			labs.GET("/:slug/articles", labController.GetLabArticles)
		}
	}

	// 需要认证的路由
	authenticated := api.Group("")
	authenticated.Use(middleware.AuthMiddleware())
	{
		// 用户信息
		authenticated.GET("/auth/profile", authController.GetProfile)

		// 文章管理
		authenticated.POST("/articles", articleController.CreateArticle)
		authenticated.PUT("/articles/:id", articleController.UpdateArticle)
		authenticated.DELETE("/articles/:id", articleController.DeleteArticle)

		// 分类管理
		authenticated.POST("/categories", categoryController.CreateCategory)
		authenticated.PUT("/categories/:id", categoryController.UpdateCategory)
		authenticated.DELETE("/categories/:id", categoryController.DeleteCategory)

		// 标签管理
		authenticated.POST("/tags", tagController.CreateTag)
		authenticated.DELETE("/tags/:id", tagController.DeleteTag)

		// 评论管理
		authenticated.PUT("/comments/:id/status", commentController.UpdateCommentStatus)
		authenticated.DELETE("/comments/:id", commentController.DeleteComment)
		authenticated.GET("/comments/pending", commentController.GetPendingComments)

		// 友情链接管理
		authenticated.POST("/links", linkController.CreateLink)
		authenticated.PUT("/links/:id", linkController.UpdateLink)
		authenticated.DELETE("/links/:id", linkController.DeleteLink)

		// 文件上传（需要认证）
		authenticated.POST("/upload/file", uploadController.UploadFile)
		authenticated.POST("/upload/image", uploadController.UploadImage)
		authenticated.DELETE("/upload/:filename", uploadController.DeleteFile)
	}

	// 管理员路由
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		// 音乐管理
		admin.POST("/music", musicController.CreateMusic)
		admin.PUT("/music/:id", musicController.UpdateMusic)
		admin.DELETE("/music/:id", musicController.DeleteMusic)
	}

	// 静态文件服务（使用配置中的路径）
	r.Static("/uploads", config.AppConfig.UploadPath)
	r.Static("/music", config.AppConfig.MusicPath)

	return r
}

