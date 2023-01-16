package models

type article struct {
	Id         int        `json:"_id"  binding:"required"`
	Title      string     `json:"title"  binding:"required"`
	Content    string     `json:"content"`
	Categories []category `json:"categories"  binding:"required"`
}

var articles []article

