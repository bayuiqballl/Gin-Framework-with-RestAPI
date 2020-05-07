package main

import (
	"github.com/bayuiqballl/Gin-framework/config"
	"github.com/bayuiqballl/Gin-framework/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// set up database
	config.InitDB()
	defer config.DB.Close()

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
