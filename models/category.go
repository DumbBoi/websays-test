package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type category struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
}

func GetCategory(name string) (category, bool) {
	files, err := ioutil.ReadDir("./category files/")
	if err != nil {
		log.Println(err)
	}
	for _, cat := range files {
		if cat.Name() == name {
			return ReadCategory(cat.Name()), true
		}
	}
	return category{}, false
}

func CreateCategory(name string, description string) (category, error) {
	cat, ok := GetCategory(name)
	if ok {
		return cat, nil
	}
	cat.Name = name
	if description != "" {
		cat.Description = description
	}

	file, _ := json.MarshalIndent(cat, "", " ")
	filepath := fmt.Sprintf("./category files/%s.json", cat.Name)
	_ = ioutil.WriteFile(filepath, file, 0644)
	return cat, nil
}

func ReadCategory(name string) category {
	var (
		result   category
		err      error
		jsonFile *os.File
	)
	filepath := fmt.Sprintf("./category files/%s.json", name)
	jsonFile, err = os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	log.Println("Successfully Opened file")
	defer jsonFile.Close()
	return result
}

func UpdateCategory(name string, description string) (category, error) {
	cat, ok := GetCategory(name)
	if ok {
		cat.Name = name
		cat.Description = description
		file, _ := json.MarshalIndent(cat, "", " ")
		filepath := fmt.Sprintf("./category files/%s.json", cat.Name)
		_ = ioutil.WriteFile(filepath, file, 0644)
		return cat, nil
	}
	return cat, nil
}

func DeleteCategory(name string) error {
	filepath := fmt.Sprintf("./category files/%s.json", name)
	err := os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}
