package main

import (
	"github.com/bayuiqballl/Gin-framework/config"
	"github.com/bayuiqballl/Gin-framework/middleware"
	"github.com/bayuiqballl/Gin-framework/routes"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	// set up database
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	// setup routing/Router
	router := gin.Default()

	// // without grouping
	// router.GET("/", getHome)
	// router.GET("/article/:title", getArticle)
	// router.POST("/articles", postArticle)

	// // Grouping endpoint
	v1 := router.Group("/api/v1/")
	{
		{
			v1.GET("/auth/:provider", routes.RedirectHandler)
			v1.GET("/auth/:provider/callback", routes.CallbackHandler)

			// testing token
			// v1.GET("/check", middleware.IsAuth(), routes.CheckToken)

			v1.GET("/profile", middleware.IsAuth(), routes.GetProfile)
			v1.GET("/article/:slug", routes.GetArticle)
			articles := v1.Group("/articles")
			{
				articles.GET("/", routes.GetHome)
				articles.GET("/tag/:tag", routes.GetArticleByTag)
				articles.GET("/id/:id", routes.GetArticleById)
				articles.POST("/", middleware.IsAuth(), routes.PostArticle)
				articles.PUT("/update/:id", middleware.IsAuth(), routes.UpdateArticle)
				articles.DELETE("/delete/:id", middleware.IsAdmin(), routes.DeleteArticle)
			}
		}
	}

	router.Run()

}
