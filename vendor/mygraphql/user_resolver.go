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
type user struct {
	id         graphqlgo.ID
	first_name string
	last_name  string
	email      string
	password   string
	token      string
	org_id     int32
	team       *team
	product    *product
	org        *org
}

// Struct for upserting
type userInput struct {
	Id        *graphqlgo.ID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Token     string
	OrgId     *int32
	Team      *teamInput
	Product   *productInput
}

// Struct for response
type userResolver struct {
	user *user
}

func ResolveUser(args struct {
	ID graphqlgo.ID
}) (response []*userResolver) {
	if args.ID != "" {
		response = append(response, &userResolver{user: MapUser(models.GetUser(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllUsers() {
		response = append(response, &userResolver{user: MapUser(val)})
	}
	return response
}

func ResolveCreateUser(args *struct {
	User *userInput
}) *userResolver {
	var user = &user{}

	if args.User.Id == nil {
		user = MapUser(models.PostUser(ReverseMapUser(args.User)))
	} else {
		user = MapUser(models.PutUser(ReverseMapUser(args.User)))
	}
	if user != nil && args.User.Team != nil {
		if args.User.Team.Id == nil {
			team := ReverseMapTeam(args.User.Team)
			if team.UserId != 0 && utils.ConvertId(user.id) != team.UserId {
				// todo throw error
				return &userResolver{}
			}
			team.UserId = utils.ConvertId(user.id)
			user.team = MapTeam(models.PostTeam(team))
		} else {
			team := ReverseMapTeam(args.User.Team)
			if team.UserId != 0 && utils.ConvertId(user.id) != team.UserId {
				// todo throw error
				return &userResolver{}
			}
			team.UserId = utils.ConvertId(user.id)
			user.team = MapTeam(models.PutTeam(team))
		}
	}
	if user != nil && args.User.Product != nil {
		if args.User.Product.Id == nil {
			product := ReverseMapProduct(args.User.Product)
			if product.UserId != 0 && utils.ConvertId(user.id) != product.UserId {
				// todo throw error
				return &userResolver{}
			}
			product.UserId = utils.ConvertId(user.id)
			user.product = MapProduct(models.PostProduct(product))
		} else {
			product := ReverseMapProduct(args.User.Product)
			if product.UserId != 0 && utils.ConvertId(user.id) != product.UserId {
				// todo throw error
				return &userResolver{}
			}
			product.UserId = utils.ConvertId(user.id)
			user.product = MapProduct(models.PutProduct(product))
		}
	}
	return &userResolver{user}
}

// For Delete
func ResolveDeleteUser(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.UserChildren) == 0 && len(models.UserInterRelation) == 0 {
		del = models.DeleteUser(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.User
		database.SQL.Model(models.User{}).Preload("Team").Preload("Product").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		if data.Team.Id != 0 {
			args.ID = utils.UintToGraphId(data.Team.Id)
			ResolveDeleteTeam(args, "")
			count++
		}

		if data.Product.Id != 0 {
			args.ID = utils.UintToGraphId(data.Product.Id)
			ResolveDeleteProduct(args, "")
			count++
		}

		del = models.DeleteUser(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.User
	database.SQL.Model(models.User{}).Preload("Team").Preload("Product").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	if data.Team.Id != 0 {
		flag++
	}

	if data.Product.Id != 0 {
		flag++
	}

	if flag == 0 {
		del = models.DeleteUser(utils.ConvertId(tempID), name)
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
func (r *userResolver) Id() graphqlgo.ID {
	return r.user.id
}
func (r *userResolver) FirstName() string {
	return r.user.first_name
}
func (r *userResolver) LastName() string {
	return r.user.last_name
}
func (r *userResolver) Email() string {
	return r.user.email
}
func (r *userResolver) Password() string {
	return r.user.password
}
func (r *userResolver) Token() string {
	return r.user.token
}
func (r *userResolver) OrgId() int32 {
	return r.user.org_id
}
func (r *userResolver) Team() *teamResolver {
	if r.user != nil {
		team := models.GetTeamOfUser(utils.ConvertId(r.user.id))
		return &teamResolver{MapTeam(team)}
	}
	return &teamResolver{r.user.team}
}
func (r *userResolver) Product() *productResolver {
	if r.user != nil {
		product := models.GetProductOfUser(utils.ConvertId(r.user.id))
		return &productResolver{MapProduct(product)}
	}
	return &productResolver{r.user.product}
}
func (r *userResolver) Org() *orgResolver {
	if r.user != nil {
		org := models.GetOrgOfUser(ReverseMap2User(r.user))
		return &orgResolver{MapOrg(org)}
	}
	return &orgResolver{r.user.org}
}

// Mapper methods
func MapUser(modelUser models.User) *user {

	if reflect.DeepEqual(modelUser, models.User{}) {
		return &user{}
	}

	// Create graphql user from models User
	user := user{
		email:      modelUser.Email,
		first_name: modelUser.FirstName,
		id:         utils.UintToGraphId(modelUser.Id),
		last_name:  modelUser.LastName,
		org_id:     int32(modelUser.OrgId),
		password:   modelUser.Password,
		token:      modelUser.Token,
	}
	return &user
}

// Reverse Mapper methods
func ReverseMapUser(mygraphqlUser *userInput) models.User {

	if reflect.DeepEqual(mygraphqlUser, userInput{}) {
		return models.User{}
	}

	// Create graphql user from models User
	var userModel models.User
	if mygraphqlUser.Id == nil {
		userModel = models.User{
			Email:     mygraphqlUser.Email,
			FirstName: mygraphqlUser.FirstName,
			LastName:  mygraphqlUser.LastName,
			Password:  mygraphqlUser.Password,
			Token:     mygraphqlUser.Token,
		}
	} else {
		userModel = models.User{
			Email:     mygraphqlUser.Email,
			FirstName: mygraphqlUser.FirstName,
			Id:        utils.ConvertId(*mygraphqlUser.Id),
			LastName:  mygraphqlUser.LastName,
			OrgId:     utils.Int32ToUint(*mygraphqlUser.OrgId),
			Password:  mygraphqlUser.Password,
			Token:     mygraphqlUser.Token,
		}
	}
	return userModel
}
func ReverseMap2User(structUser *user) models.User {

	if reflect.DeepEqual(structUser, user{}) {
		return models.User{}
	}

	// Create graphql user from models User
	modelUser := models.User{
		Email:     structUser.email,
		FirstName: structUser.first_name,
		Id:        utils.ConvertId(structUser.id),
		LastName:  structUser.last_name,
		OrgId:     uint(structUser.org_id),
		Password:  structUser.password,
		Token:     structUser.token,
	}
	return modelUser
}
