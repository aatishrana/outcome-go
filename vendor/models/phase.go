package models

import (
	database "database"
	generator "generator"
)

type Phase struct {
	Id      uint     `gorm:"column:id" json:"id,omitempty"`
	Name    string   `gorm:"column:name" json:"name,omitempty"`
	Sprints []Sprint `json:"Sprints,omitempty"`
}

func (Phase) TableName() string {
	return "phase"
}

// Child entities
var PhaseChildren = []string{}

// Inter entities
var PhaseInterRelation = []generator.InterEntity{
	generator.InterEntity{TableName: "sprint_phase", StructName: "SprintPhase"},
}

// This method will return a list of all Phases
func GetAllPhases(limit int, offset int) []Phase {
	data := []Phase{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
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
func GetPhasesOfSprint(sprintid uint) []Phase {
	data := []Phase{}
	data2 := []SprintPhase{}
	database.SQL.Debug().Where("sprint_id = ?", sprintid).Find(&data2)
	var sliceOfId []uint
	for _, v := range data2 {
		sliceOfId = append(sliceOfId, v.PhaseId)
	}
	database.SQL.Debug().Where("id IN (?)", sliceOfId).Find(&data)
	return data
}
