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
			v1.GET("/check", middleware.IsAuth(), routes.CheckToken)

			articles := v1.Group("/article")
			{
				articles.GET("/", routes.GetHome)
				articles.GET("/:slug", routes.GetArticle)
				articles.POST("/", routes.PostArticle)

			}
		}
	}

	router.Run()

}
