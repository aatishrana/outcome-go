package models

import (
	database "database"
	generator "generator"
)

type User struct {
	Id              uint             `gorm:"column:id" json:"id,omitempty"`
	FirstName       string           `gorm:"column:first_name" json:"first_name,omitempty"`
	LastName        string           `gorm:"column:last_name" json:"last_name,omitempty"`
	Email           string           `gorm:"column:email" json:"email,omitempty"`
	Password        string           `gorm:"column:password" json:"password,omitempty"`
	Token           string           `gorm:"column:token" json:"token,omitempty"`
	OrgId           uint             `gorm:"column:org_id" json:"org_id,omitempty"`
	Teams           []Team           `json:"Teams,omitempty"`
	UserTeams       []UserTeam       `gorm:"ForeignKey:user_id;AssociationForeignKey:id" json:"UserTeams,omitempty"`
	Team            Team             `gorm:"ForeignKey:user_id;AssociationForeignKey:id" json:"Team,omitempty"`
	Product         Product          `gorm:"ForeignKey:user_id;AssociationForeignKey:id" json:"Product,omitempty"`
	ProductBackLogs []ProductBackLog `gorm:"ForeignKey:user_id;AssociationForeignKey:id" json:"ProductBackLogs,omitempty"`
	Project         Project          `gorm:"ForeignKey:user_id;AssociationForeignKey:id" json:"Project,omitempty"`
	Tasks           []Task           `gorm:"ForeignKey:user_id;AssociationForeignKey:id" json:"Tasks,omitempty"`
	Org             Org              `gorm:"ForeignKey:OrgId" json:"Org,omitempty"`
}

func (User) TableName() string {
	return "user"
}

// Child entities
var UserChildren = []string{"Team", "Product", "ProductBackLogs", "Project", "Tasks"}

// Inter entities
var UserInterRelation = []generator.InterEntity{
	generator.InterEntity{TableName: "user_team", StructName: "UserTeam"},
}

// This method will return a list of all Users
func GetAllUsers() []User {
	data := []User{}
	database.SQL.Find(&data)
	return data
}

// This method will return one User based on id
func GetUser(ID uint) User {
	data := User{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one User in db
func PostUser(data User) User {
	database.SQL.Create(&data)
	return data
}

// This method will update User based on id
func PutUser(newData User) User {
	oldData := User{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetUser(newData.Id)
}

// This method will delete User based on id
func DeleteUser(ID uint, parent string) bool {
	var data User
	var del bool
	if parent == "" {
		database.SQL.Where("user.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("user.id=(?)", ID).Where("user.user_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
func GetUsersOfTeam(teamid uint) []User {
	data := []User{}
	data2 := []UserTeam{}
	database.SQL.Debug().Where("team_id = ?", teamid).Find(&data2)
	var sliceOfId []uint
	for _, v := range data2 {
		sliceOfId = append(sliceOfId, v.UserId)
	}
	database.SQL.Debug().Where("id IN (?)", sliceOfId).Find(&data)
	return data
}
func GetUserOfTeam(team Team) User {
	data := User{}
	database.SQL.Debug().Where("id = ?", team.UserId).Find(&data)
	return data
}
func GetUserOfProduct(product Product) User {
	data := User{}
	database.SQL.Debug().Where("id = ?", product.UserId).Find(&data)
	return data
}
func GetUserOfProductBackLog(productbacklog ProductBackLog) User {
	data := User{}
	database.SQL.Debug().Where("id = ?", productbacklog.UserId).Find(&data)
	return data
}
func GetUserOfProject(project Project) User {
	data := User{}
	database.SQL.Debug().Where("id = ?", project.UserId).Find(&data)
	return data
}
func GetUserOfTask(task Task) User {
	data := User{}
	database.SQL.Debug().Where("id = ?", task.UserId).Find(&data)
	return data
}
func GetUsersOfOrg(orgid uint) []User {
	data := Org{}
	database.SQL.Debug().Preload("Users").Where("id = ?", orgid).Find(&data)
	return data.Users
}
