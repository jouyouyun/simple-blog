package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	router *gin.Engine
)

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initRoutes()
	router.Run(":8080")
}

func initRoutes() {
	router.GET("/", showArticles)
	router.GET("/article/view/:id", showArticle)
}

func showArticles(c *gin.Context) {
	articles := getArticles()
	// Call the HTML method to render a template
	c.HTML(
		http.StatusOK,
		"index.html",
		// Pass the data to page
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func showArticle(c *gin.Context) {
	s := c.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid url")
		return
	}

	article, err := getArticle(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Not found")
		return
	}

	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
	)
}
