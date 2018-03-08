package models

import (
	database "database"
	generator "generator"
)

type Org struct {
	Id       uint      `gorm:"column:id" json:"id,omitempty"`
	Name     string    `gorm:"column:name" json:"name,omitempty"`
	Users    []User    `gorm:"ForeignKey:org_id;AssociationForeignKey:id" json:"Users,omitempty"`
	Teams    []Team    `gorm:"ForeignKey:org_id;AssociationForeignKey:id" json:"Teams,omitempty"`
	Products []Product `gorm:"ForeignKey:org_id;AssociationForeignKey:id" json:"Products,omitempty"`
}

func (Org) TableName() string {
	return "org"
}

// Child entities
var OrgChildren = []string{"Users", "Teams", "Products"}

// Inter entities
var OrgInterRelation = []generator.InterEntity{}

// This method will return a list of all Orgs
func GetAllOrgs(limit int, offset int) []Org {
	data := []Org{}
	database.SQL.Limit(limit).Offset(offset).Find(&data)
	return data
}

// This method will return one Org based on id
func GetOrg(ID uint) Org {
	data := Org{}
	database.SQL.First(&data, ID)
	return data
}

// This method will insert one Org in db
func PostOrg(data Org) Org {
	database.SQL.Create(&data)
	return data
}

// This method will update Org based on id
func PutOrg(newData Org) Org {
	oldData := Org{Id: newData.Id}
	database.SQL.Model(&oldData).Updates(newData)
	return GetOrg(newData.Id)
}

// This method will delete Org based on id
func DeleteOrg(ID uint, parent string) bool {
	var data Org
	var del bool
	if parent == "" {
		database.SQL.Where("org.id=(?)", ID).First(&data)
	} else {
		database.SQL.Where("org.id=(?)", ID).Where("org.org_type=(?)", parent).First(&data)
	}
	if data.Id != 0 {
		database.SQL.Delete(&data)
		del = true
	}
	return del
}
func GetOrgOfUser(user User) Org {
	data := Org{}
	database.SQL.Debug().Where("id = ?", user.OrgId).Find(&data)
	return data
}
func GetOrgOfTeam(team Team) Org {
	data := Org{}
	database.SQL.Debug().Where("id = ?", team.OrgId).Find(&data)
	return data
}
func GetOrgOfProduct(product Product) Org {
	data := Org{}
	database.SQL.Debug().Where("id = ?", product.OrgId).Find(&data)
	return data
}
