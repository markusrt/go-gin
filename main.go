// main.go

package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	// Handle Index
	router.GET("/", showIndexPage)
	router.GET("/articles", showIndexPage)
	// Handle GET requests at /articles/some_article_id
	router.GET("/articles/:article_id", getArticle)

	router.Run()

}
