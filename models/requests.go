package models

type CreateArticleRequest struct {
	Title    string `binding:"required"`
	Content  string
	Category []string `binding:"required"`
}
