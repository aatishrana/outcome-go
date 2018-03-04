package models

import (
	database "database"
	generator "generator"
)

type UserTeam struct {
	Id     uint `gorm:"column:id" json:"id,omitempty"`
	UserId uint `gorm:"column:user_id" json:"user_id,omitempty"`
	TeamId uint `gorm:"column:team_id" json:"team_id,omitempty"`
}

func (UserTeam) TableName() string {
	return "user_team"
}

// Child entities
var UserTeamChildren = []string{}

// Inter entities
var UserTeamInterRelation = []generator.InterEntity{}

// This method will return a list of all UserTeams
func GetAllUserTeams() []UserTeam {
	data := []UserTeam{}
	database.SQL.Find(&data)
	return data
}

// This method will return one UserTeam based on id
func GetUserTeam(ID uint) UserTeam {
	data := UserTeam{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one UserTeam in db
func PostUserTeam(data UserTeam) UserTeam {
	database.SQL.Create(&data)
	return data
}

// This method will update UserTeam based on id
func PutUserTeam(newData UserTeam) UserTeam {
	oldData := UserTeam{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetUserTeam(newData.Id)
}

// This method will delete UserTeam based on id
func DeleteUserTeam(ID uint, parent string) bool {
	var data UserTeam
	var del bool
	if parent == "" {
		database.SQL.Where("user_team.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("user_team.id=(?)", ID).Where("user_team.userteam_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
