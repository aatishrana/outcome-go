package mygraphql

import (
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
	org     *org
	user    *user
}

// Struct for upserting
type teamInput struct {
	Id     *graphqlgo.ID
	Name   string
	UserId *int32
	OrgId  *int32
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
	for _, val := range models.GetAllTeams() {
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
func (r *teamResolver) Org() *orgResolver {
	if r.team != nil {
		org := models.GetOrgOfTeam(ReverseMap2Team(r.team))
		return &orgResolver{MapOrg(org)}
	}
	return &orgResolver{r.team.org}
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
