package models

import (
	database "database"
	generator "generator"
)

type Project struct {
	Id        uint     `gorm:"column:id" json:"id,omitempty"`
	Name      string   `gorm:"column:name" json:"name,omitempty"`
	UserId    uint     `gorm:"column:user_id" json:"user_id,omitempty"`
	TeamId    uint     `gorm:"column:team_id" json:"team_id,omitempty"`
	ProductId uint     `gorm:"column:product_id" json:"product_id,omitempty"`
	Storys    []Story  `gorm:"ForeignKey:project_id;AssociationForeignKey:id" json:"Storys,omitempty"`
	Sprints   []Sprint `gorm:"ForeignKey:project_id;AssociationForeignKey:id" json:"Sprints,omitempty"`
	Product   Product  `gorm:"ForeignKey:ProductId" json:"Product,omitempty"`
}

func (Project) TableName() string {
	return "project"
}

// Child entities
var ProjectChildren = []string{"Storys", "Sprints"}

// Inter entities
var ProjectInterRelation = []generator.InterEntity{}

// This method will return a list of all Projects
func GetAllProjects(limit int, offset int) []Project {
	data := []Project{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
	return data
}

// This method will return one Project based on id
func GetProject(ID uint) Project {
	data := Project{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Project in db
func PostProject(data Project) Project {
	database.SQL.Create(&data)
	return data
}

// This method will update Project based on id
func PutProject(newData Project) Project {
	oldData := Project{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetProject(newData.Id)
}

// This method will delete Project based on id
func DeleteProject(ID uint, parent string) bool {
	var data Project
	var del bool
	if parent == "" {
		database.SQL.Where("project.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("project.id=(?)", ID).Where("project.project_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
func GetProjectOfStory(story Story) Project {
	data := Project{}
	database.SQL.Debug().Where("id = ?", story.ProjectId).Find(&data)
	return data
}
func GetProjectOfSprint(sprint Sprint) Project {
	data := Project{}
	database.SQL.Debug().Where("id = ?", sprint.ProjectId).Find(&data)
	return data
}
func GetProjectOfUser(userid uint) Project {
	data := User{}
	database.SQL.Debug().Preload("Project").Where("id = ?", userid).Find(&data)
	return data.Project
}
func GetProjectOfTeam(teamid uint) Project {
	data := Team{}
	database.SQL.Debug().Preload("Project").Where("id = ?", teamid).Find(&data)
	return data.Project
}
func GetProjectsOfProduct(productid uint) []Project {
	data := Product{}
	database.SQL.Debug().Preload("Projects").Where("id = ?", productid).Find(&data)
	return data.Projects
}
