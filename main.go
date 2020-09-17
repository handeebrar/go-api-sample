package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	model "./models"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	page := model.Page{ID: 3, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/categories"}
	categories := loadCategories()
	products := loadProducts()
	productMappings := loadProductMappings()

	var newCategories []model.Category

	for _, category := range categories {

		for _, productMapping := range productMappings {
			if category.ID == productMapping.CategoryID {
				for _, product := range products {
					if productMapping.ProductID == product.ID {
						category.Products = append(category.Products, product)
					}
				}
			}
		}
		newCategories = append(newCategories, category)
	}

	viewModel := model.CategoryViewModel{Page: page, Categories: newCategories}

	t, _ := template.ParseFiles("template/page.html")
	t.Execute(w, viewModel)
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadCategories() []model.Category {
	bytes, _ := ioutil.ReadFile("json/categories.json")
	var categories []model.Category
	json.Unmarshal(bytes, &categories)
	return categories
}

func loadProducts() []model.Product {
	bytes, _ := ioutil.ReadFile("json/products.json")
	var products []model.Product
	json.Unmarshal(bytes, &products)
	return products
}

func loadProductMappings() []model.ProductsMapping {
	bytes, _ := ioutil.ReadFile("json/categoryProductsMappings.json")
	var productMappings []model.ProductsMapping
	json.Unmarshal(bytes, &productMappings)
	return productMappings
}
