package mygraphql

import (
	database "database"
	fmt "fmt"
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type org struct {
	id       graphqlgo.ID
	name     string
	users    []*user
	teams    []*team
	products []*product
}

// Struct for upserting
type orgInput struct {
	Id       *graphqlgo.ID
	Name     string
	Users    *[]userInput
	Teams    *[]teamInput
	Products *[]productInput
}

// Struct for response
type orgResolver struct {
	org *org
}

func ResolveOrg(args struct {
	ID graphqlgo.ID
}) (response []*orgResolver) {
	if args.ID != "" {
		response = append(response, &orgResolver{org: MapOrg(models.GetOrg(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllOrgs() {
		response = append(response, &orgResolver{org: MapOrg(val)})
	}
	return response
}

func ResolveCreateOrg(args *struct {
	Org *orgInput
}) *orgResolver {
	var org = &org{}

	if args.Org.Id == nil {
		org = MapOrg(models.PostOrg(ReverseMapOrg(args.Org)))
	} else {
		org = MapOrg(models.PutOrg(ReverseMapOrg(args.Org)))
	}
	if org != nil && args.Org.Users != nil {
		for _, dev := range *args.Org.Users {
			if dev.Id == nil {
				user := ReverseMapUser(&dev)
				if user.OrgId != 0 && utils.ConvertId(org.id) != user.OrgId {
					// todo throw error
					return &orgResolver{}
				}
				user.OrgId = utils.ConvertId(org.id)
				org.users = append(org.users, MapUser(models.PostUser(user)))
			} else {
				user := ReverseMapUser(&dev)
				if user.OrgId != 0 && utils.ConvertId(org.id) != user.OrgId {
					// todo throw error
					return &orgResolver{}
				}
				user.OrgId = utils.ConvertId(org.id)
				org.users = append(org.users, MapUser(models.PutUser(user)))
			}
		}
	}
	if org != nil && args.Org.Teams != nil {
		for _, dev := range *args.Org.Teams {
			if dev.Id == nil {
				team := ReverseMapTeam(&dev)
				if team.OrgId != 0 && utils.ConvertId(org.id) != team.OrgId {
					// todo throw error
					return &orgResolver{}
				}
				team.OrgId = utils.ConvertId(org.id)
				org.teams = append(org.teams, MapTeam(models.PostTeam(team)))
			} else {
				team := ReverseMapTeam(&dev)
				if team.OrgId != 0 && utils.ConvertId(org.id) != team.OrgId {
					// todo throw error
					return &orgResolver{}
				}
				team.OrgId = utils.ConvertId(org.id)
				org.teams = append(org.teams, MapTeam(models.PutTeam(team)))
			}
		}
	}
	if org != nil && args.Org.Products != nil {
		for _, dev := range *args.Org.Products {
			if dev.Id == nil {
				product := ReverseMapProduct(&dev)
				if product.OrgId != 0 && utils.ConvertId(org.id) != product.OrgId {
					// todo throw error
					return &orgResolver{}
				}
				product.OrgId = utils.ConvertId(org.id)
				org.products = append(org.products, MapProduct(models.PostProduct(product)))
			} else {
				product := ReverseMapProduct(&dev)
				if product.OrgId != 0 && utils.ConvertId(org.id) != product.OrgId {
					// todo throw error
					return &orgResolver{}
				}
				product.OrgId = utils.ConvertId(org.id)
				org.products = append(org.products, MapProduct(models.PutProduct(product)))
			}
		}
	}
	return &orgResolver{org}
}

// For Delete
func ResolveDeleteOrg(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.OrgChildren) == 0 && len(models.OrgInterRelation) == 0 {
		del = models.DeleteOrg(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.Org
		database.SQL.Model(models.Org{}).Preload("Users").Preload("Teams").Preload("Products").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		for _, v := range data.Users {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteUser(args, "")
			count++
		}

		for _, v := range data.Teams {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteTeam(args, "")
			count++
		}

		for _, v := range data.Products {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteProduct(args, "")
			count++
		}

		del = models.DeleteOrg(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.Org
	database.SQL.Model(models.Org{}).Preload("Users").Preload("Teams").Preload("Products").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	for _, v := range data.Users {
		if v.Id != 0 {
			flag++
		}
	}

	for _, v := range data.Teams {
		if v.Id != 0 {
			flag++
		}
	}

	for _, v := range data.Products {
		if v.Id != 0 {
			flag++
		}
	}

	if flag == 0 {
		del = models.DeleteOrg(utils.ConvertId(tempID), name)
		count++
		response = &count
	} else {
		// show error
		fmt.Println("Cannot Delete :", tempID)
		del = false
		response = &count
	}
	return response
}

// Fields resolvers
func (r *orgResolver) Id() graphqlgo.ID {
	return r.org.id
}
func (r *orgResolver) Name() string {
	return r.org.name
}
func (r *orgResolver) Users() []*userResolver {
	var users []*userResolver
	if r.org != nil {
		user := models.GetUsersOfOrg(utils.ConvertId(r.org.id))
		for _, value := range user {
			users = append(users, &userResolver{MapUser(value)})
		}
		return users
	}
	for _, value := range r.org.users {
		users = append(users, &userResolver{value})
	}
	return users
}
func (r *orgResolver) Teams() []*teamResolver {
	var teams []*teamResolver
	if r.org != nil {
		team := models.GetTeamsOfOrg(utils.ConvertId(r.org.id))
		for _, value := range team {
			teams = append(teams, &teamResolver{MapTeam(value)})
		}
		return teams
	}
	for _, value := range r.org.teams {
		teams = append(teams, &teamResolver{value})
	}
	return teams
}
func (r *orgResolver) Products() []*productResolver {
	var products []*productResolver
	if r.org != nil {
		product := models.GetProductsOfOrg(utils.ConvertId(r.org.id))
		for _, value := range product {
			products = append(products, &productResolver{MapProduct(value)})
		}
		return products
	}
	for _, value := range r.org.products {
		products = append(products, &productResolver{value})
	}
	return products
}

// Mapper methods
func MapOrg(modelOrg models.Org) *org {

	if reflect.DeepEqual(modelOrg, models.Org{}) {
		return &org{}
	}

	// Create graphql org from models Org
	org := org{
		id:   utils.UintToGraphId(modelOrg.Id),
		name: modelOrg.Name,
	}
	return &org
}

// Reverse Mapper methods
func ReverseMapOrg(mygraphqlOrg *orgInput) models.Org {

	if reflect.DeepEqual(mygraphqlOrg, orgInput{}) {
		return models.Org{}
	}

	// Create graphql org from models Org
	var orgModel models.Org
	if mygraphqlOrg.Id == nil {
		orgModel = models.Org{Name: mygraphqlOrg.Name}
	} else {
		orgModel = models.Org{
			Id:   utils.ConvertId(*mygraphqlOrg.Id),
			Name: mygraphqlOrg.Name,
		}
	}
	return orgModel
}
func ReverseMap2Org(structOrg *org) models.Org {

	if reflect.DeepEqual(structOrg, org{}) {
		return models.Org{}
	}

	// Create graphql org from models Org
	modelOrg := models.Org{
		Id:   utils.ConvertId(structOrg.id),
		Name: structOrg.name,
	}
	return modelOrg
}
