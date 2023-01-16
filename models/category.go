package models

type category struct {
	Id   int    `binding:"required"`
	Name string `binding:"required"`
}

var categories []category

func GetCategory(name string) (category, bool) {
	for _, cat := range categories {
		if cat.Name == name {
			return cat, true
		}
	}
	return category{}, false
}
