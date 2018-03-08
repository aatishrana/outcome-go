package models

import (
	database "database"
	generator "generator"
)

type Story struct {
	Id               uint           `gorm:"column:id" json:"id,omitempty"`
	Desc             string         `gorm:"column:desc" json:"desc,omitempty"`
	Status           string         `gorm:"column:status" json:"status,omitempty"`
	Point            uint           `gorm:"column:point" json:"point,omitempty"`
	ProductBackLogId uint           `gorm:"column:product_back_log_id" json:"product_back_log_id,omitempty"`
	ProjectId        uint           `gorm:"column:project_id" json:"project_id,omitempty"`
	SprintId         uint           `gorm:"column:sprint_id" json:"sprint_id,omitempty"`
	Tasks            []Task         `gorm:"ForeignKey:story_id;AssociationForeignKey:id" json:"Tasks,omitempty"`
	ProductBackLog   ProductBackLog `gorm:"ForeignKey:ProductBackLogId" json:"ProductBackLog,omitempty"`
	Project          Project        `gorm:"ForeignKey:ProjectId" json:"Project,omitempty"`
	Sprint           Sprint         `gorm:"ForeignKey:SprintId" json:"Sprint,omitempty"`
}

func (Story) TableName() string {
	return "story"
}

// Child entities
var StoryChildren = []string{"Tasks"}

// Inter entities
var StoryInterRelation = []generator.InterEntity{}

// This method will return a list of all Storys
func GetAllStorys(limit int, offset int) []Story {
	data := []Story{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
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
func GetStoryOfTask(task Task) Story {
	data := Story{}
	database.SQL.Debug().Where("id = ?", task.StoryId).Find(&data)
	return data
}
func GetStorysOfProductBackLog(productbacklogid uint) []Story {
	data := ProductBackLog{}
	database.SQL.Debug().Preload("Storys").Where("id = ?", productbacklogid).Find(&data)
	return data.Storys
}
func GetStorysOfProject(projectid uint) []Story {
	data := Project{}
	database.SQL.Debug().Preload("Storys").Where("id = ?", projectid).Find(&data)
	return data.Storys
}
func GetStorysOfSprint(sprintid uint) []Story {
	data := Sprint{}
	database.SQL.Debug().Preload("Storys").Where("id = ?", sprintid).Find(&data)
	return data.Storys
}
