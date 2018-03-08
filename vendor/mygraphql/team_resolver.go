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
type team struct {
	id      graphqlgo.ID
	name    string
	user_id int32
	org_id  int32
	project *project
	org     *org
	users   []*user
	user    *user
}

// Struct for upserting
type teamInput struct {
	Id      *graphqlgo.ID
	Name    string
	UserId  *int32
	OrgId   *int32
	Project *projectInput
}

// Struct for response
type teamResolver struct {
	team *team
}

func ResolveTeam(args struct {
	ID graphqlgo.ID
}) (response []*teamResolver) {
	if args.ID != "" {
		response = append(response, &teamResolver{team: MapTeam(models.GetTeam(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllTeams(utils.GetDefaultLimitOffset()) {
		response = append(response, &teamResolver{team: MapTeam(val)})
	}
	return response
}

func ResolveCreateTeam(args *struct {
	Team *teamInput
}) *teamResolver {
	var team = &team{}

	if args.Team.Id == nil {
		team = MapTeam(models.PostTeam(ReverseMapTeam(args.Team)))
	} else {
		team = MapTeam(models.PutTeam(ReverseMapTeam(args.Team)))
	}
	if team != nil && args.Team.Project != nil {
		if args.Team.Project.Id == nil {
			project := ReverseMapProject(args.Team.Project)
			if project.TeamId != 0 && utils.ConvertId(team.id) != project.TeamId {
				// todo throw error
				return &teamResolver{}
			}
			project.TeamId = utils.ConvertId(team.id)
			team.project = MapProject(models.PostProject(project))
		} else {
			project := ReverseMapProject(args.Team.Project)
			if project.TeamId != 0 && utils.ConvertId(team.id) != project.TeamId {
				// todo throw error
				return &teamResolver{}
			}
			project.TeamId = utils.ConvertId(team.id)
			team.project = MapProject(models.PutProject(project))
		}
	}
	return &teamResolver{team}
}

// For Delete
func ResolveDeleteTeam(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.TeamChildren) == 0 && len(models.TeamInterRelation) == 0 {
		del = models.DeleteTeam(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.Team
		database.SQL.Model(models.Team{}).Preload("Project").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		if data.Project.Id != 0 {
			args.ID = utils.UintToGraphId(data.Project.Id)
			ResolveDeleteProject(args, "")
			count++
		}

		del = models.DeleteTeam(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.Team
	database.SQL.Model(models.Team{}).Preload("Project").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	if data.Project.Id != 0 {
		flag++
	}

	if flag == 0 {
		del = models.DeleteTeam(utils.ConvertId(tempID), name)
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
func (r *teamResolver) Id() graphqlgo.ID {
	return r.team.id
}
func (r *teamResolver) Name() string {
	return r.team.name
}
func (r *teamResolver) UserId() int32 {
	return r.team.user_id
}
func (r *teamResolver) OrgId() int32 {
	return r.team.org_id
}
func (r *teamResolver) Project() *projectResolver {
	if r.team != nil {
		project := models.GetProjectOfTeam(utils.ConvertId(r.team.id))
		return &projectResolver{MapProject(project)}
	}
	return &projectResolver{r.team.project}
}
func (r *teamResolver) Org() *orgResolver {
	if r.team != nil {
		org := models.GetOrgOfTeam(ReverseMap2Team(r.team))
		return &orgResolver{MapOrg(org)}
	}
	return &orgResolver{r.team.org}
}
func (r *teamResolver) Users() []*userResolver {
	var users []*userResolver
	if r.team != nil {
		user := models.GetUsersOfTeam(utils.ConvertId(r.team.id))
		for _, value := range user {
			users = append(users, &userResolver{MapUser(value)})
		}
		return users
	}
	for _, value := range r.team.users {
		users = append(users, &userResolver{value})
	}
	return users
}
func (r *teamResolver) User() *userResolver {
	if r.team != nil {
		user := models.GetUserOfTeam(ReverseMap2Team(r.team))
		return &userResolver{MapUser(user)}
	}
	return &userResolver{r.team.user}
}

// Mapper methods
func MapTeam(modelTeam models.Team) *team {

	if reflect.DeepEqual(modelTeam, models.Team{}) {
		return &team{}
	}

	// Create graphql team from models Team
	team := team{
		id:      utils.UintToGraphId(modelTeam.Id),
		name:    modelTeam.Name,
		org_id:  int32(modelTeam.OrgId),
		user_id: int32(modelTeam.UserId),
	}
	return &team
}

// Reverse Mapper methods
func ReverseMapTeam(mygraphqlTeam *teamInput) models.Team {

	if reflect.DeepEqual(mygraphqlTeam, teamInput{}) {
		return models.Team{}
	}

	// Create graphql team from models Team
	var teamModel models.Team
	if mygraphqlTeam.Id == nil {
		teamModel = models.Team{Name: mygraphqlTeam.Name}
	} else {
		teamModel = models.Team{
			Id:     utils.ConvertId(*mygraphqlTeam.Id),
			Name:   mygraphqlTeam.Name,
			OrgId:  utils.Int32ToUint(*mygraphqlTeam.OrgId),
			UserId: utils.Int32ToUint(*mygraphqlTeam.UserId),
		}
	}
	return teamModel
}
func ReverseMap2Team(structTeam *team) models.Team {

	if reflect.DeepEqual(structTeam, team{}) {
		return models.Team{}
	}

	// Create graphql team from models Team
	modelTeam := models.Team{
		Id:     utils.ConvertId(structTeam.id),
		Name:   structTeam.name,
		OrgId:  uint(structTeam.org_id),
		UserId: uint(structTeam.user_id),
	}
	return modelTeam
}
