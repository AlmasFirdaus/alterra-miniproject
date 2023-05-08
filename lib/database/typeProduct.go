package database

import (
	"alterra-miniproject/config"
	"alterra-miniproject/models"
)

func GetTypeProducts() (interface{}, error) {
	var typeProducts []models.TypeProduct

	if e := config.DB.Find(&typeProducts).Error; e != nil {
		return nil, e
	}
	return typeProducts, nil
}

func GetTypeProduct(id int) (interface{}, error) {
	var typeProduct []models.TypeProduct

	if e := config.DB.First(&typeProduct, id).Error; e != nil {
		return nil, e
	}

	return typeProduct, nil
}

func CreateTypeProduct(typeProduct models.TypeProduct) (interface{}, error) {
	if e := config.DB.Create(&typeProduct).Error; e != nil {
		return nil, e
	}

	return typeProduct, nil
}

func UpdateTypeProduct(tp models.TypeProduct) (interface{}, error) {
	var typeProductUpdate models.TypeProduct
	config.DB.First(&typeProductUpdate, tp.ID)

	if e := config.DB.Model(&tp).Updates(models.TypeProduct{Name: tp.Name}).Error; e != nil {
		return nil, e
	}
	return typeProductUpdate, nil
}

func DeleteTypeProduct(id int) (interface{}, error) {
	var typeProduct models.TypeProduct
	if e := config.DB.Delete(&typeProduct, id).Error; e != nil {
		return nil, e
	}
	return typeProduct, nil
}
