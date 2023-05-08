package database

import (
	"alterra-miniproject/config"
	"alterra-miniproject/models"
)

func GetProducts() (interface{}, error) {
	var products []models.Product

	if e := config.DB.Preload("TypeProduct").Find(&products).Error; e != nil {
		return nil, e
	}
	return products, nil
}

func GetProduct(id int) (interface{}, error) {
	var product []models.Product

	if e := config.DB.Preload("TypeProduct").First(&product, id).Error; e != nil {
		return nil, e
	}
	return product, nil
}

func CreateProduct(product models.Product) (interface{}, error) {
	if e := config.DB.Preload("TypeProduct").Create(&product).Error; e != nil {
		return nil, e
	}

	return product, nil
}

func UpdateProduct(product models.Product) (interface{}, error) {
	var productUpdate models.Product
	config.DB.First(&productUpdate, product.ID)

	if e := config.DB.Model(&productUpdate).Updates(models.Product{Merk: product.Merk, Weight: product.Weight, Qty: product.Qty, TypeProductID: product.TypeProductID}).Error; e != nil {
		return nil, e
	}
	return productUpdate, nil
}

func DeleteProduct(id int) (interface{}, error) {
	var product models.Product
	if e := config.DB.Delete(&product, id).Error; e != nil {
		return nil, e
	}
	return product, nil
}
