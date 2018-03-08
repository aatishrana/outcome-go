package models

import (
	database "database"
	generator "generator"
)

type Team struct {
	Id      uint    `gorm:"column:id" json:"id,omitempty"`
	Name    string  `gorm:"column:name" json:"name,omitempty"`
	UserId  uint    `gorm:"column:user_id" json:"user_id,omitempty"`
	OrgId   uint    `gorm:"column:org_id" json:"org_id,omitempty"`
	Project Project `gorm:"ForeignKey:team_id;AssociationForeignKey:id" json:"Project,omitempty"`
	Org     Org     `gorm:"ForeignKey:OrgId" json:"Org,omitempty"`
	Users   []User  `json:"Users,omitempty"`
}

func (Team) TableName() string {
	return "team"
}

// Child entities
var TeamChildren = []string{"Project"}

// Inter entities
var TeamInterRelation = []generator.InterEntity{
	generator.InterEntity{TableName: "user_team", StructName: "UserTeam"},
}

// This method will return a list of all Teams
func GetAllTeams(limit int, offset int) []Team {
	data := []Team{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
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
func GetTeamOfProject(project Project) Team {
	data := Team{}
	database.SQL.Debug().Where("id = ?", project.TeamId).Find(&data)
	return data
}
func GetTeamsOfOrg(orgid uint) []Team {
	data := Org{}
	database.SQL.Debug().Preload("Teams").Where("id = ?", orgid).Find(&data)
	return data.Teams
}
func GetTeamsOfUser(userid uint) []Team {
	data := []Team{}
	data2 := []UserTeam{}
	database.SQL.Debug().Where("user_id = ?", userid).Find(&data2)
	var sliceOfId []uint
	for _, v := range data2 {
		sliceOfId = append(sliceOfId, v.TeamId)
	}
	database.SQL.Debug().Where("id IN (?)", sliceOfId).Find(&data)
	return data
}
func GetTeamOfUser(userid uint) Team {
	data := User{}
	database.SQL.Debug().Preload("Team").Where("id = ?", userid).Find(&data)
	return data.Team
}
