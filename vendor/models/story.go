package models

import (
	database "database"
	generator "generator"
)

type Story struct {
	Id               uint   `gorm:"column:id" json:"id,omitempty"`
	Desc             string `gorm:"column:desc" json:"desc,omitempty"`
	Status           string `gorm:"column:status" json:"status,omitempty"`
	Point            uint   `gorm:"column:point" json:"point,omitempty"`
	ProductBackLogId uint   `gorm:"column:product_back_log_id" json:"product_back_log_id,omitempty"`
	ProjectId        uint   `gorm:"column:project_id" json:"project_id,omitempty"`
	SprintId         uint   `gorm:"column:sprint_id" json:"sprint_id,omitempty"`
}

func (Story) TableName() string {
	return "story"
}

// Child entities
var StoryChildren = []string{}

// Inter entities
var StoryInterRelation = []generator.InterEntity{}

// This method will return a list of all Storys
func GetAllStorys() []Story {
	data := []Story{}
	database.SQL.Find(&data)
	return data
}

// This method will return one Story based on id
func GetStory(ID uint) Story {
	data := Story{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Story in db
func PostStory(data Story) Story {
	database.SQL.Create(&data)
	return data
}

// This method will update Story based on id
func PutStory(newData Story) Story {
	oldData := Story{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetStory(newData.Id)
}

// This method will delete Story based on id
func DeleteStory(ID uint, parent string) bool {
	var data Story
	var del bool
	if parent == "" {
		database.SQL.Where("story.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("story.id=(?)", ID).Where("story.story_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
