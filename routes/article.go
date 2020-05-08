package routes

import (
	"strconv"
	"time"

	"github.com/bayuiqballl/Gin-framework/config"
	"github.com/bayuiqballl/Gin-framework/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetProfile(c *gin.Context) {
	var user models.User
	user_id := int(c.MustGet("jwt_user_id").(float64))

	item := config.DB.Where("id = ?", user_id).Preload("Articles", "user_id = ?", user_id).Find(&user)

	c.JSON(200, gin.H{
		"status": "berhasil ke halaman profile",
		"data":   item,
	})

}

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

	var oldItem models.Article
	slug := slug.Make(c.PostForm("title"))

	if !config.DB.First(&oldItem, "slug = ?", slug).RecordNotFound() {
		// generate string baru
		slug = slug + "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	item := models.Article{
		Title:  c.PostForm("title"),
		Desc:   c.PostForm("desc"),
		Tag:    c.PostForm("tag"),
		Slug:   slug,
		UserID: uint(c.MustGet("jwt_user_id").(float64)),
	}

	// kalau slug sama, maka generate random slug
	//

	config.DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "post berhasil",
		"data":   item,
	})
}

func GetArticleByTag(c *gin.Context) {
	tag := c.Param("tag")
	items := []models.Article{}

	config.DB.Where("tag LIKE ?", "%"+tag+"%").Find(&items)

	c.JSON(200, gin.H{"dara": items})

}

func UpdateArticle(c *gin.Context) {

	// ambil data detail dari database/API
	// mengolah
	id := c.Param("id")

	var item models.Article

	if config.DB.First(&item, "id = ?", id).RecordNotFound() {
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort()
		return
	}

	if uint(c.MustGet("jwt_user_id").(float64)) != item.UserID {
		c.JSON(403, gin.H{"status": "error", "message": "this data is forbidden"})
		c.Abort()
		return
	}

	config.DB.Model(&item).Where("id = ?", id).Updates(models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Tag:   c.PostForm("tag"),
	})

	c.JSON(200, gin.H{
		"status": "berhasil update",
		"data":   item,
	})
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	config.DB.Where("id = ?", id).Delete(&article)
	c.JSON(200, gin.H{
		"status": "berhasil delete",
		"data":   article,
	})
}
