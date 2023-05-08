package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Merk          string `json:"merk" form:"merk"`
	Weight        int    `json:"weight" form:"weight"`
	Qty           int    `json:"qty" form:"qty"`
	TypeProductID int    `json:"type_product_id" form:"type_product_id"`
	TypeProduct   TypeProduct
}
