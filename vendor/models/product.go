package models

import (
	database "database"
	generator "generator"
)

type Product struct {
	Id              uint             `gorm:"column:id" json:"id,omitempty"`
	Name            string           `gorm:"column:name" json:"name,omitempty"`
	Desc            string           `gorm:"column:desc" json:"desc,omitempty"`
	UserId          uint             `gorm:"column:user_id" json:"user_id,omitempty"`
	OrgId           uint             `gorm:"column:org_id" json:"org_id,omitempty"`
	ProductBackLogs []ProductBackLog `gorm:"ForeignKey:product_id;AssociationForeignKey:id" json:"ProductBackLogs,omitempty"`
	Projects        []Project        `gorm:"ForeignKey:product_id;AssociationForeignKey:id" json:"Projects,omitempty"`
	Org             Org              `gorm:"ForeignKey:OrgId" json:"Org,omitempty"`
}

func (Product) TableName() string {
	return "product"
}

// Child entities
var ProductChildren = []string{"ProductBackLogs", "Projects"}

// Inter entities
var ProductInterRelation = []generator.InterEntity{}

// This method will return a list of all Products
func GetAllProducts(limit int, offset int) []Product {
	data := []Product{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
	return data
}

// This method will return one Product based on id
func GetProduct(ID uint) Product {
	data := Product{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Product in db
func PostProduct(data Product) Product {
	database.SQL.Create(&data)
	return data
}

// This method will update Product based on id
func PutProduct(newData Product) Product {
	oldData := Product{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetProduct(newData.Id)
}

// This method will delete Product based on id
func DeleteProduct(ID uint, parent string) bool {
	var data Product
	var del bool
	if parent == "" {
		database.SQL.Where("product.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("product.id=(?)", ID).Where("product.product_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
func GetProductOfProductBackLog(productbacklog ProductBackLog) Product {
	data := Product{}
	database.SQL.Debug().Where("id = ?", productbacklog.ProductId).Find(&data)
	return data
}
func GetProductOfProject(project Project) Product {
	data := Product{}
	database.SQL.Debug().Where("id = ?", project.ProductId).Find(&data)
	return data
}
func GetProductsOfOrg(orgid uint) []Product {
	data := Org{}
	database.SQL.Debug().Preload("Products").Where("id = ?", orgid).Find(&data)
	return data.Products
}
func GetProductOfUser(userid uint) Product {
	data := User{}
	database.SQL.Debug().Preload("Product").Where("id = ?", userid).Find(&data)
	return data.Product
}
