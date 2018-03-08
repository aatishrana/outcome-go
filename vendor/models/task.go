package models

import (
	database "database"
	generator "generator"
)

type Task struct {
	Id            uint   `gorm:"column:id" json:"id,omitempty"`
	SprintId      uint   `gorm:"column:sprint_id" json:"sprint_id,omitempty"`
	StoryId       uint   `gorm:"column:story_id" json:"story_id,omitempty"`
	SprintPhaseId uint   `gorm:"column:sprint_phase_id" json:"sprint_phase_id,omitempty"`
	UserId        uint   `gorm:"column:user_id" json:"user_id,omitempty"`
	Point         uint   `gorm:"column:point" json:"point,omitempty"`
	StartDtTm     string `gorm:"column:start_dt_tm" json:"start_dt_tm,omitempty"`
	EndDtTm       string `gorm:"column:end_dt_tm" json:"end_dt_tm,omitempty"`
	User          User   `gorm:"ForeignKey:UserId" json:"User,omitempty"`
	Story         Story  `gorm:"ForeignKey:StoryId" json:"Story,omitempty"`
	Sprint        Sprint `gorm:"ForeignKey:SprintId" json:"Sprint,omitempty"`
}

func (Task) TableName() string {
	return "task"
}

// Child entities
var TaskChildren = []string{}

// Inter entities
var TaskInterRelation = []generator.InterEntity{}

// This method will return a list of all Tasks
func GetAllTasks(limit int, offset int) []Task {
	data := []Task{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
	return data
}

// This method will return one Task based on id
func GetTask(ID uint) Task {
	data := Task{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Task in db
func PostTask(data Task) Task {
	database.SQL.Create(&data)
	return data
}

// This method will update Task based on id
func PutTask(newData Task) Task {
	oldData := Task{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetTask(newData.Id)
}

// This method will delete Task based on id
func DeleteTask(ID uint, parent string) bool {
	var data Task
	var del bool
	if parent == "" {
		database.SQL.Where("task.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("task.id=(?)", ID).Where("task.task_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
func GetTasksOfUser(userid uint) []Task {
	data := User{}
	database.SQL.Debug().Preload("Tasks").Where("id = ?", userid).Find(&data)
	return data.Tasks
}
func GetTasksOfStory(storyid uint) []Task {
	data := Story{}
	database.SQL.Debug().Preload("Tasks").Where("id = ?", storyid).Find(&data)
	return data.Tasks
}
func GetTasksOfSprint(sprintid uint) []Task {
	data := Sprint{}
	database.SQL.Debug().Preload("Tasks").Where("id = ?", sprintid).Find(&data)
	return data.Tasks
}
