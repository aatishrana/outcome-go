package models

import (
	database "database"
	generator "generator"
)

type Team struct {
	Id     uint   `gorm:"column:id" json:"id,omitempty"`
	Name   string `gorm:"column:name" json:"name,omitempty"`
	UserId uint   `gorm:"column:user_id" json:"user_id,omitempty"`
	OrgId  uint   `gorm:"column:org_id" json:"org_id,omitempty"`
	Org    Org    `gorm:"ForeignKey:OrgId" json:"Org,omitempty"`
}

func (Team) TableName() string {
	return "team"
}

// Child entities
var TeamChildren = []string{}

// Inter entities
var TeamInterRelation = []generator.InterEntity{}

// This method will return a list of all Teams
func GetAllTeams() []Team {
	data := []Team{}
	database.SQL.Find(&data)
	return data
}

// This method will return one Team based on id
func GetTeam(ID uint) Team {
	data := Team{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Team in db
func PostTeam(data Team) Team {
	database.SQL.Create(&data)
	return data
}

// This method will update Team based on id
func PutTeam(newData Team) Team {
	oldData := Team{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetTeam(newData.Id)
}

// This method will delete Team based on id
func DeleteTeam(ID uint, parent string) bool {
	var data Team
	var del bool
	if parent == "" {
		database.SQL.Where("team.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("team.id=(?)", ID).Where("team.team_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
func GetTeamsOfOrg(orgid uint) []Team {
	data := Org{}
	database.SQL.Debug().Preload("Teams").Where("id = ?", orgid).Find(&data)
	return data.Teams
}
func GetTeamOfUser(userid uint) Team {
	data := User{}
	database.SQL.Debug().Preload("Team").Where("id = ?", userid).Find(&data)
	return data.Team
}
