package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(
		c,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		}, "index.html")
}

func getArticle(c *gin.Context) {
	//Checking Article ID Valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		//Check Article Exist
		if article, err := getArticleByID(articleID); err == nil {
			//Call HTML
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
