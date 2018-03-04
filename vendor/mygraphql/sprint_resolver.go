package mygraphql

import (
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type sprint struct {
	id         graphqlgo.ID
	name       string
	start_dt   string
	end_dt     string
	project_id int32
}

// Struct for upserting
type sprintInput struct {
	Id        *graphqlgo.ID
	Name      string
	StartDt   string
	EndDt     string
	ProjectId *int32
}

// Struct for response
type sprintResolver struct {
	sprint *sprint
}

func ResolveSprint(args struct {
	ID graphqlgo.ID
}) (response []*sprintResolver) {
	if args.ID != "" {
		response = append(response, &sprintResolver{sprint: MapSprint(models.GetSprint(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllSprints() {
		response = append(response, &sprintResolver{sprint: MapSprint(val)})
	}
	return response
}

func ResolveCreateSprint(args *struct {
	Sprint *sprintInput
}) *sprintResolver {
	var sprint = &sprint{}

	if args.Sprint.Id == nil {
		sprint = MapSprint(models.PostSprint(ReverseMapSprint(args.Sprint)))
	} else {
		sprint = MapSprint(models.PutSprint(ReverseMapSprint(args.Sprint)))
	}
	return &sprintResolver{sprint}
}

// For Delete
func ResolveDeleteSprint(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.SprintChildren) == 0 && len(models.SprintInterRelation) == 0 {
		del = models.DeleteSprint(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	return response
}

// Fields resolvers
func (r *sprintResolver) Id() graphqlgo.ID {
	return r.sprint.id
}
func (r *sprintResolver) Name() string {
	return r.sprint.name
}
func (r *sprintResolver) StartDt() string {
	return r.sprint.start_dt
}
func (r *sprintResolver) EndDt() string {
	return r.sprint.end_dt
}
func (r *sprintResolver) ProjectId() int32 {
	return r.sprint.project_id
}

// Mapper methods
func MapSprint(modelSprint models.Sprint) *sprint {

	if reflect.DeepEqual(modelSprint, models.Sprint{}) {
		return &sprint{}
	}

	// Create graphql sprint from models Sprint
	sprint := sprint{
		end_dt:     modelSprint.EndDt,
		id:         utils.UintToGraphId(modelSprint.Id),
		name:       modelSprint.Name,
		project_id: int32(modelSprint.ProjectId),
		start_dt:   modelSprint.StartDt,
	}
	return &sprint
}

// Reverse Mapper methods
func ReverseMapSprint(mygraphqlSprint *sprintInput) models.Sprint {

	if reflect.DeepEqual(mygraphqlSprint, sprintInput{}) {
		return models.Sprint{}
	}

	// Create graphql sprint from models Sprint
	var sprintModel models.Sprint
	if mygraphqlSprint.Id == nil {
		sprintModel = models.Sprint{
			EndDt:   mygraphqlSprint.EndDt,
			Name:    mygraphqlSprint.Name,
			StartDt: mygraphqlSprint.StartDt,
		}
	} else {
		sprintModel = models.Sprint{
			EndDt:     mygraphqlSprint.EndDt,
			Id:        utils.ConvertId(*mygraphqlSprint.Id),
			Name:      mygraphqlSprint.Name,
			ProjectId: utils.Int32ToUint(*mygraphqlSprint.ProjectId),
			StartDt:   mygraphqlSprint.StartDt,
		}
	}
	return sprintModel
}
func ReverseMap2Sprint(structSprint *sprint) models.Sprint {

	if reflect.DeepEqual(structSprint, sprint{}) {
		return models.Sprint{}
	}

	// Create graphql sprint from models Sprint
	modelSprint := models.Sprint{
		EndDt:     structSprint.end_dt,
		Id:        utils.ConvertId(structSprint.id),
		Name:      structSprint.name,
		ProjectId: uint(structSprint.project_id),
		StartDt:   structSprint.start_dt,
	}
	return modelSprint
}
