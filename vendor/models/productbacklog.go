package models

import (
	database "database"
	generator "generator"
)

type ProductBackLog struct {
	Id        uint   `gorm:"column:id" json:"id,omitempty"`
	Desc      string `gorm:"column:desc" json:"desc,omitempty"`
	TypeCd    string `gorm:"column:type_cd" json:"type_cd,omitempty"`
	Priority  string `gorm:"column:priority" json:"priority,omitempty"`
	UserId    uint   `gorm:"column:user_id" json:"user_id,omitempty"`
	ProductId uint   `gorm:"column:product_id" json:"product_id,omitempty"`
}

func (ProductBackLog) TableName() string {
	return "product_back_log"
}

// Child entities
var ProductBackLogChildren = []string{}

// Inter entities
var ProductBackLogInterRelation = []generator.InterEntity{}

// This method will return a list of all ProductBackLogs
func GetAllProductBackLogs() []ProductBackLog {
	data := []ProductBackLog{}
	database.SQL.Find(&data)
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
