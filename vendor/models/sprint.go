package models

import (
	database "database"
	generator "generator"
)

type Sprint struct {
	Id        uint   `gorm:"column:id" json:"id,omitempty"`
	Name      string `gorm:"column:name" json:"name,omitempty"`
	StartDt   string `gorm:"column:start_dt" json:"start_dt,omitempty"`
	EndDt     string `gorm:"column:end_dt" json:"end_dt,omitempty"`
	ProjectId uint   `gorm:"column:project_id" json:"project_id,omitempty"`
}

func (Sprint) TableName() string {
	return "sprint"
}

// Child entities
var SprintChildren = []string{}

// Inter entities
var SprintInterRelation = []generator.InterEntity{}

// This method will return a list of all Sprints
func GetAllSprints() []Sprint {
	data := []Sprint{}
	database.SQL.Find(&data)
	return data
}

// This method will return one Sprint based on id
func GetSprint(ID uint) Sprint {
	data := Sprint{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Sprint in db
func PostSprint(data Sprint) Sprint {
	database.SQL.Create(&data)
	return data
}

// This method will update Sprint based on id
func PutSprint(newData Sprint) Sprint {
	oldData := Sprint{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetSprint(newData.Id)
}

// This method will delete Sprint based on id
func DeleteSprint(ID uint, parent string) bool {
	var data Sprint
	var del bool
	if parent == "" {
		database.SQL.Where("sprint.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("sprint.id=(?)", ID).Where("sprint.sprint_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
