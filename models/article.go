package models

type article struct {
	Id         int        `json:"_id"  binding:"required"`
	Title      string     `json:"title"  binding:"required"`
	Content    string     `json:"content"`
	Categories []category `json:"categories"  binding:"required"`
}

var articles []article

func CreateArticle(request CreateArticleRequest) (article, error) {
	var art article
	art.Id = len(articles)
	art.Title = request.Title
	art.Content = request.Content
	for _, catstr := range request.Category {
		cat, ok := GetCategory(catstr)
		if !ok {
			cat, _ = CreateCategory(catstr)
		}
		art.Categories = append(art.Categories, cat)
	}
	articles = append(articles, art)
	return art, nil
}

func WriteAritcle() {

}

func UpdateAritcle() {

}

func DeleteAritcle() {

}
