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
	id              graphqlgo.ID
	first_name      string
	last_name       string
	email           string
	password        string
	token           string
	org_id          int32
	teams           []*team
	team            *team
	product         *product
	productbacklogs []*productbacklog
	project         *project
	tasks           []*task
	org             *org
}

// Struct for upserting
type userInput struct {
	Id              *graphqlgo.ID
	FirstName       *string
	LastName        *string
	Email           *string
	Password        *string
	Token           *string
	OrgId           *int32
	Teams           *[]teamInput
	Team            *teamInput
	Product         *productInput
	ProductBackLogs *[]productbacklogInput
	Project         *projectInput
	Tasks           *[]taskInput
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
	for _, val := range models.GetAllUsers(utils.GetDefaultLimitOffset()) {
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
	if user != nil && args.User.Teams != nil {
		for _, dev := range *args.User.Teams {
			if dev.Id == nil {
				team := ReverseMapTeam(&dev)
				user.teams = append(user.teams, MapTeam(models.PostTeam(team)))

				var data = userteamInput{}
				userteam := ReverseMapUserTeam(&data)
				userId := utils.ConvertId(user.id)
				var teamId uint
				for _, val := range user.teams {
					teamId = utils.ConvertId(val.id)
				}
				userteam.UserId = userId
				userteam.TeamId = teamId
				models.PostUserTeam(userteam)

			} else {
				team := ReverseMapTeam(&dev)
				user.teams = append(user.teams, MapTeam(models.PutTeam(team)))

				var data = userteamInput{}
				userteam := ReverseMapUserTeam(&data)
				userId := utils.ConvertId(user.id)
				var teamId uint
				for _, val := range user.teams {
					teamId = utils.ConvertId(val.id)
				}
				userteam.UserId = userId
				userteam.TeamId = teamId
				models.PostUserTeam(userteam)

			}
		}
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
	if user != nil && args.User.ProductBackLogs != nil {
		for _, dev := range *args.User.ProductBackLogs {
			if dev.Id == nil {
				productbacklog := ReverseMapProductBackLog(&dev)
				if productbacklog.UserId != 0 && utils.ConvertId(user.id) != productbacklog.UserId {
					// todo throw error
					return &userResolver{}
				}
				productbacklog.UserId = utils.ConvertId(user.id)
				user.productbacklogs = append(user.productbacklogs, MapProductBackLog(models.PostProductBackLog(productbacklog)))
			} else {
				productbacklog := ReverseMapProductBackLog(&dev)
				if productbacklog.UserId != 0 && utils.ConvertId(user.id) != productbacklog.UserId {
					// todo throw error
					return &userResolver{}
				}
				productbacklog.UserId = utils.ConvertId(user.id)
				user.productbacklogs = append(user.productbacklogs, MapProductBackLog(models.PutProductBackLog(productbacklog)))
			}
		}
	}
	if user != nil && args.User.Project != nil {
		if args.User.Project.Id == nil {
			project := ReverseMapProject(args.User.Project)
			if project.UserId != 0 && utils.ConvertId(user.id) != project.UserId {
				// todo throw error
				return &userResolver{}
			}
			project.UserId = utils.ConvertId(user.id)
			user.project = MapProject(models.PostProject(project))
		} else {
			project := ReverseMapProject(args.User.Project)
			if project.UserId != 0 && utils.ConvertId(user.id) != project.UserId {
				// todo throw error
				return &userResolver{}
			}
			project.UserId = utils.ConvertId(user.id)
			user.project = MapProject(models.PutProject(project))
		}
	}
	if user != nil && args.User.Tasks != nil {
		for _, dev := range *args.User.Tasks {
			if dev.Id == nil {
				task := ReverseMapTask(&dev)
				if task.UserId != 0 && utils.ConvertId(user.id) != task.UserId {
					// todo throw error
					return &userResolver{}
				}
				task.UserId = utils.ConvertId(user.id)
				user.tasks = append(user.tasks, MapTask(models.PostTask(task)))
			} else {
				task := ReverseMapTask(&dev)
				if task.UserId != 0 && utils.ConvertId(user.id) != task.UserId {
					// todo throw error
					return &userResolver{}
				}
				task.UserId = utils.ConvertId(user.id)
				user.tasks = append(user.tasks, MapTask(models.PutTask(task)))
			}
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
		database.SQL.Model(models.User{}).Preload("UserTeams").Preload("Team").Preload("Product").Preload("ProductBackLogs").Preload("Project").Preload("Tasks").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		for _, v := range data.UserTeams {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteUserTeam(args, "")
			count++
		}

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

		for _, v := range data.ProductBackLogs {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteProductBackLog(args, "")
			count++
		}

		if data.Project.Id != 0 {
			args.ID = utils.UintToGraphId(data.Project.Id)
			ResolveDeleteProject(args, "")
			count++
		}

		for _, v := range data.Tasks {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteTask(args, "")
			count++
		}

		del = models.DeleteUser(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.User
	database.SQL.Model(models.User{}).Preload("UserTeams").Preload("Team").Preload("Product").Preload("ProductBackLogs").Preload("Project").Preload("Tasks").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	for _, v := range data.UserTeams {
		if v.Id != 0 {
			flag++
		}
	}

	if data.Team.Id != 0 {
		flag++
	}

	if data.Product.Id != 0 {
		flag++
	}

	for _, v := range data.ProductBackLogs {
		if v.Id != 0 {
			flag++
		}
	}

	if data.Project.Id != 0 {
		flag++
	}

	for _, v := range data.Tasks {
		if v.Id != 0 {
			flag++
		}
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
func (r *userResolver) Teams() []*teamResolver {
	var teams []*teamResolver
	if r.user != nil {
		team := models.GetTeamsOfUser(utils.ConvertId(r.user.id))
		for _, value := range team {
			teams = append(teams, &teamResolver{MapTeam(value)})
		}
		return teams
	}
	for _, value := range r.user.teams {
		teams = append(teams, &teamResolver{value})
	}
	return teams
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
func (r *userResolver) ProductBackLogs() []*productbacklogResolver {
	var productbacklogs []*productbacklogResolver
	if r.user != nil {
		productbacklog := models.GetProductBackLogsOfUser(utils.ConvertId(r.user.id))
		for _, value := range productbacklog {
			productbacklogs = append(productbacklogs, &productbacklogResolver{MapProductBackLog(value)})
		}
		return productbacklogs
	}
	for _, value := range r.user.productbacklogs {
		productbacklogs = append(productbacklogs, &productbacklogResolver{value})
	}
	return productbacklogs
}
func (r *userResolver) Project() *projectResolver {
	if r.user != nil {
		project := models.GetProjectOfUser(utils.ConvertId(r.user.id))
		return &projectResolver{MapProject(project)}
	}
	return &projectResolver{r.user.project}
}
func (r *userResolver) Tasks() []*taskResolver {
	var tasks []*taskResolver
	if r.user != nil {
		task := models.GetTasksOfUser(utils.ConvertId(r.user.id))
		for _, value := range task {
			tasks = append(tasks, &taskResolver{MapTask(value)})
		}
		return tasks
	}
	for _, value := range r.user.tasks {
		tasks = append(tasks, &taskResolver{value})
	}
	return tasks
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

	var id uint = 0;
	var orgId uint = 0;
	email := "";
	firstName := "";
	lastName := "";
	password := "";
	token := "";

	if mygraphqlUser.Id != nil {
		id = utils.ConvertId(*mygraphqlUser.Id)
	}
	if mygraphqlUser.OrgId != nil {
		orgId = utils.Int32ToUint(*mygraphqlUser.OrgId)
	}
	if mygraphqlUser.Email != nil {
		email = *mygraphqlUser.Email
	}
	if mygraphqlUser.FirstName != nil {
		firstName = *mygraphqlUser.FirstName
	}
	if mygraphqlUser.LastName != nil {
		lastName = *mygraphqlUser.LastName
	}
	if mygraphqlUser.Password != nil {
		password = *mygraphqlUser.Password
	}
	if mygraphqlUser.Token != nil {
		token = *mygraphqlUser.Token
	}

	return models.User{
		Email:     email,
		FirstName: firstName,
		Id:        id,
		LastName:  lastName,
		OrgId:     orgId,
		Password:  password,
		Token:     token,
	}
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
