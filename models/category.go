package models

type category struct {
	Id   int    `binding:"required"`
	Name string `binding:"required"`
}

var categories []category

func CreateCategory(name string) (category, error) {
	cat, ok := GetCategory(name)
	if !ok {
		cat.Id = len(categories)
		cat.Name = name
	}
	categories = append(categories, cat)
	return cat, nil
}

func GetCategory(name string) (category, bool) {
	for _, cat := range categories {
		if cat.Name == name {
			return cat, true
		}
	}
	return category{}, false
}
