package models

import (
	database "database"
	generator "generator"
)

type SprintPhase struct {
	Id       uint `gorm:"column:id" json:"id,omitempty"`
	SprintId uint `gorm:"column:sprint_id" json:"sprint_id,omitempty"`
	PhaseId  uint `gorm:"column:phase_id" json:"phase_id,omitempty"`
}

func (SprintPhase) TableName() string {
	return "sprint_phase"
}

// Child entities
var SprintPhaseChildren = []string{}

// Inter entities
var SprintPhaseInterRelation = []generator.InterEntity{}

// This method will return a list of all SprintPhases
func GetAllSprintPhases() []SprintPhase {
	data := []SprintPhase{}
	database.SQL.Find(&data)
	return data
}

// This method will return one SprintPhase based on id
func GetSprintPhase(ID uint) SprintPhase {
	data := SprintPhase{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one SprintPhase in db
func PostSprintPhase(data SprintPhase) SprintPhase {
	var oldData []SprintPhase
	database.SQL.Find(&oldData)
	for _, val := range oldData {
		if val.SprintId == data.SprintId && val.PhaseId == data.PhaseId {
			return SprintPhase{}
		}
	}
	database.SQL.Create(&data)
	return data
}

// This method will update SprintPhase based on id
func PutSprintPhase(newData SprintPhase) SprintPhase {
	oldData := SprintPhase{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetSprintPhase(newData.Id)
}

// This method will delete SprintPhase based on id
func DeleteSprintPhase(ID uint, parent string) bool {
	var data SprintPhase
	var del bool
	if parent == "" {
		database.SQL.Where("sprint_phase.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("sprint_phase.id=(?)", ID).Where("sprint_phase.sprintphase_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
