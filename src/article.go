package main

import (
	"fmt"
)

type articleInfo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type articleInfos []*articleInfo

var articleList = articleInfos{
	&articleInfo{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	&articleInfo{ID: 2, Title: "Article 2", Content: "Article 2 body"},
	&articleInfo{ID: 3, Title: "Article 3", Content: "Article 3 body"},
}

func getArticles() articleInfos {
	return articleList
}

func getArticle(id int) (*articleInfo, error) {
	info := getArticles().Get(id)
	if info == nil {
		return nil, fmt.Errorf("Invalid article id: %v", id)
	}
	return info, nil
}

func (infos articleInfos) Get(id int) *articleInfo {
	for _, info := range infos {
		if info.ID == id {
			return info
		}
	}
	return nil
}
