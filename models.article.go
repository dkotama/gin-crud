package main

import (
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Bombing assault in afghanistan", Content: "Lorem Ipsum Dolor sit amor quiwlk hauytu"},
	article{ID: 2, Title: "June was Awesome", Content: "Lorem Ipsum Dolor sit amor quiwlk hauytu"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	// pass kalo gak ketemu sama errornya
	return nil, errors.New("Article Not Found")
}
