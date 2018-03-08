package mygraphql

import (
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type userteam struct {
	id      graphqlgo.ID
	user_id int32
	team_id int32
}

// Struct for upserting
type userteamInput struct {
	Id     *graphqlgo.ID
	UserId *int32
	TeamId *int32
}

// Struct for response
type userteamResolver struct {
	userteam *userteam
}

func ResolveUserTeam(args struct {
	ID graphqlgo.ID
}) (response []*userteamResolver) {
	if args.ID != "" {
		response = append(response, &userteamResolver{userteam: MapUserTeam(models.GetUserTeam(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllUserTeams(utils.GetDefaultLimitOffset()) {
		response = append(response, &userteamResolver{userteam: MapUserTeam(val)})
	}
	return response
}

func ResolveCreateUserTeam(args *struct {
	UserTeam *userteamInput
}) *userteamResolver {
	var userteam = &userteam{}

	if args.UserTeam.Id == nil {
		userteam = MapUserTeam(models.PostUserTeam(ReverseMapUserTeam(args.UserTeam)))
	} else {
		userteam = MapUserTeam(models.PutUserTeam(ReverseMapUserTeam(args.UserTeam)))
	}
	return &userteamResolver{userteam}
}

// For Delete
func ResolveDeleteUserTeam(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.UserTeamChildren) == 0 && len(models.UserTeamInterRelation) == 0 {
		del = models.DeleteUserTeam(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	return response
}

// Fields resolvers
func (r *userteamResolver) Id() graphqlgo.ID {
	return r.userteam.id
}
func (r *userteamResolver) UserId() int32 {
	return r.userteam.user_id
}
func (r *userteamResolver) TeamId() int32 {
	return r.userteam.team_id
}

// Mapper methods
func MapUserTeam(modelUserTeam models.UserTeam) *userteam {

	if reflect.DeepEqual(modelUserTeam, models.UserTeam{}) {
		return &userteam{}
	}

	// Create graphql userteam from models UserTeam
	userteam := userteam{
		id:      utils.UintToGraphId(modelUserTeam.Id),
		team_id: int32(modelUserTeam.TeamId),
		user_id: int32(modelUserTeam.UserId),
	}
	return &userteam
}

// Reverse Mapper methods
func ReverseMapUserTeam(mygraphqlUserTeam *userteamInput) models.UserTeam {

	if reflect.DeepEqual(mygraphqlUserTeam, userteamInput{}) {
		return models.UserTeam{}
	}

	// Create graphql userteam from models UserTeam
	var userteamModel models.UserTeam
	if mygraphqlUserTeam.Id == nil {
		userteamModel = models.UserTeam{}
	} else {
		userteamModel = models.UserTeam{
			Id:     utils.ConvertId(*mygraphqlUserTeam.Id),
			TeamId: utils.Int32ToUint(*mygraphqlUserTeam.TeamId),
			UserId: utils.Int32ToUint(*mygraphqlUserTeam.UserId),
		}
	}
	return userteamModel
}
func ReverseMap2UserTeam(structUserTeam *userteam) models.UserTeam {

	if reflect.DeepEqual(structUserTeam, userteam{}) {
		return models.UserTeam{}
	}

	// Create graphql userteam from models UserTeam
	modelUserTeam := models.UserTeam{
		Id:     utils.ConvertId(structUserTeam.id),
		TeamId: uint(structUserTeam.team_id),
		UserId: uint(structUserTeam.user_id),
	}
	return modelUserTeam
}
