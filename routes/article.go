package routes

import (
	"github.com/bayuiqballl/Gin-framework/config"
	"github.com/bayuiqballl/Gin-framework/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetHome(c *gin.Context) {
	items := []models.Article{}
	config.DB.Find(&items)
	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   items,
	})
}

//  GET
func GetArticle(c *gin.Context) {

	// ambil data detail dari database/API
	// mengolah
	slug := c.Param("slug")

	var item models.Article

	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"status": "berhasil",
		"data":   item,
	})
}

// POST
func PostArticle(c *gin.Context) {
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}

	// kalau slug sama, maka generate random slug
	//

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "post berhasil",
		"data":   item,
	})
}
