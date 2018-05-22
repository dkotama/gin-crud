package main

func initializeRoutes() {
	//Handle index route
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
}
