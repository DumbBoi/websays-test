package models

import (
	"encoding/json"
	"log"

	"github.com/DumbBoi/websays-test/dbs"
)

type article struct {
	Title      string     `json:"title"  binding:"required"`
	Content    string     `json:"content"`
	Categories []category `json:"categories"  binding:"required"`
}
type ArticleRequest struct {
	Title    string `binding:"required"`
	Content  string
	Category []string `binding:"required"`
}

var redisCleint = dbs.GetRedisObject()

func articleRequestToArticle(request ArticleRequest) article {
	var art article
	art.Title = request.Title
	art.Content = request.Content
	for _, catstr := range request.Category {
		cat, ok := GetCategory(catstr)
		if !ok {
			cat, _ = CreateCategory(catstr)
		}
		art.Categories = append(art.Categories, cat)
	}
	return art
}

func CreateArticle(request ArticleRequest) (article, error) {
	art := articleRequestToArticle(request)
	dbs.RedisWrite(art.Title, art, redisCleint)
	return art, nil
}

func ReadAritcle(key string) article {
	var articleRequestObj ArticleRequest
	obj, err := dbs.RedisRead(key, redisCleint)
	if err != nil {
		log.Println(err.Error())
	}
	json.Unmarshal([]byte(obj), &articleRequestObj)
	art := articleRequestToArticle(articleRequestObj)
	return art
}

func UpdateAritcle() {

}

func DeleteAritcle() {

}
