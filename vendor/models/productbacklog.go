package models

import (
	database "database"
	generator "generator"
)

type ProductBackLog struct {
	Id        uint    `gorm:"column:id" json:"id,omitempty"`
	Desc      string  `gorm:"column:desc" json:"desc,omitempty"`
	TypeCd    string  `gorm:"column:type_cd" json:"type_cd,omitempty"`
	Priority  string  `gorm:"column:priority" json:"priority,omitempty"`
	UserId    uint    `gorm:"column:user_id" json:"user_id,omitempty"`
	ProductId uint    `gorm:"column:product_id" json:"product_id,omitempty"`
	Storys    []Story `gorm:"ForeignKey:product_back_log_id;AssociationForeignKey:id" json:"Storys,omitempty"`
	User      User    `gorm:"ForeignKey:UserId" json:"User,omitempty"`
	Product   Product `gorm:"ForeignKey:ProductId" json:"Product,omitempty"`
}

func (ProductBackLog) TableName() string {
	return "product_back_log"
}

// Child entities
var ProductBackLogChildren = []string{"Storys"}

// Inter entities
var ProductBackLogInterRelation = []generator.InterEntity{}

// This method will return a list of all ProductBackLogs
func GetAllProductBackLogs(limit int, offset int) []ProductBackLog {
	data := []ProductBackLog{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
	return data
}

// This method will return one ProductBackLog based on id
func GetProductBackLog(ID uint) ProductBackLog {
	data := ProductBackLog{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one ProductBackLog in db
func PostProductBackLog(data ProductBackLog) ProductBackLog {
	database.SQL.Create(&data)
	return data
}

// This method will update ProductBackLog based on id
func PutProductBackLog(newData ProductBackLog) ProductBackLog {
	oldData := ProductBackLog{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetProductBackLog(newData.Id)
}

// This method will delete ProductBackLog based on id
func DeleteProductBackLog(ID uint, parent string) bool {
	var data ProductBackLog
	var del bool
	if parent == "" {
		database.SQL.Where("product_back_log.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("product_back_log.id=(?)", ID).Where("product_back_log.productbacklog_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
func GetProductBackLogOfStory(story Story) ProductBackLog {
	data := ProductBackLog{}
	database.SQL.Debug().Where("id = ?", story.ProductBackLogId).Find(&data)
	return data
}
func GetProductBackLogsOfUser(userid uint) []ProductBackLog {
	data := User{}
	database.SQL.Debug().Preload("ProductBackLogs").Where("id = ?", userid).Find(&data)
	return data.ProductBackLogs
}
func GetProductBackLogsOfProduct(productid uint) []ProductBackLog {
	data := Product{}
	database.SQL.Debug().Preload("ProductBackLogs").Where("id = ?", productid).Find(&data)
	return data.ProductBackLogs
}
