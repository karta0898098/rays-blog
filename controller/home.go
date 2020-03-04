package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"rays-blog/database"
	"rays-blog/model"
)

func Home(c *gin.Context) {

	var tags []model.Tag
	var articles []model.Article
	var recentArticles []model.Article

	database.Context.Model(model.Tag{}).Find(&tags)
	database.Context.Model(model.Article{}).Limit(5).Find(&recentArticles)

	//filter Query String tags
	if tag := c.Query("tags"); tag != "" {
		//filter
		database.Context.Model(model.Article{}).Where("Tag IN (?) ", tag).Find(&articles)
	} else {
		//find all articles
		database.Context.Model(model.Article{}).Find(&articles)
	}

	viewData := gin.H{
		"tags":           tags,
		"articles":       articles,
		"recentArticles": recentArticles,
	}

	c.HTML(http.StatusOK, "index.html", viewData)
}

func Article(c *gin.Context) {

	var tags []model.Tag
	var recentArticles []model.Article
	var article model.Article

	database.Context.Model(model.Tag{}).Find(&tags)
	database.Context.Model(model.Article{}).Limit(5).Find(&recentArticles)

	err := database.Context.Model(model.Article{}).Where("Id = ?", c.Param("id")).Find(&article).Error

	if gorm.IsRecordNotFoundError(err) {
		article = model.Article{
			Title:   "Not Found",
			Content: "",
		}
	}

	viewData := gin.H{
		"tags":           tags,
		"article":        article,
		"recentArticles": recentArticles,
	}

	c.HTML(http.StatusOK, "article.html", viewData)
}
