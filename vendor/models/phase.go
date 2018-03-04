package models

import (
	database "database"
	generator "generator"
)

type Phase struct {
	Id   uint   `gorm:"column:id" json:"id,omitempty"`
	Name string `gorm:"column:name" json:"name,omitempty"`
}

func (Phase) TableName() string {
	return "phase"
}

// Child entities
var PhaseChildren = []string{}

// Inter entities
var PhaseInterRelation = []generator.InterEntity{}

// This method will return a list of all Phases
func GetAllPhases() []Phase {
	data := []Phase{}
	database.SQL.Find(&data)
	return data
}

// This method will return one Phase based on id
func GetPhase(ID uint) Phase {
	data := Phase{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Phase in db
func PostPhase(data Phase) Phase {
	database.SQL.Create(&data)
	return data
}

// This method will update Phase based on id
func PutPhase(newData Phase) Phase {
	oldData := Phase{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetPhase(newData.Id)
}

// This method will delete Phase based on id
func DeletePhase(ID uint, parent string) bool {
	var data Phase
	var del bool
	if parent == "" {
		database.SQL.Where("phase.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("phase.id=(?)", ID).Where("phase.phase_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
