package mygraphql

import (
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type sprintphase struct {
	id        graphqlgo.ID
	sprint_id int32
	phase_id  int32
}

// Struct for upserting
type sprintphaseInput struct {
	Id       *graphqlgo.ID
	SprintId *int32
	PhaseId  *int32
}

// Struct for response
type sprintphaseResolver struct {
	sprintphase *sprintphase
}

func ResolveSprintPhase(args struct {
	ID graphqlgo.ID
}) (response []*sprintphaseResolver) {
	if args.ID != "" {
		response = append(response, &sprintphaseResolver{sprintphase: MapSprintPhase(models.GetSprintPhase(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllSprintPhases() {
		response = append(response, &sprintphaseResolver{sprintphase: MapSprintPhase(val)})
	}
	return response
}

func ResolveCreateSprintPhase(args *struct {
	SprintPhase *sprintphaseInput
}) *sprintphaseResolver {
	var sprintphase = &sprintphase{}

	if args.SprintPhase.Id == nil {
		sprintphase = MapSprintPhase(models.PostSprintPhase(ReverseMapSprintPhase(args.SprintPhase)))
	} else {
		sprintphase = MapSprintPhase(models.PutSprintPhase(ReverseMapSprintPhase(args.SprintPhase)))
	}
	return &sprintphaseResolver{sprintphase}
}

// For Delete
func ResolveDeleteSprintPhase(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.SprintPhaseChildren) == 0 && len(models.SprintPhaseInterRelation) == 0 {
		del = models.DeleteSprintPhase(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	return response
}

// Fields resolvers
func (r *sprintphaseResolver) Id() graphqlgo.ID {
	return r.sprintphase.id
}
func (r *sprintphaseResolver) SprintId() int32 {
	return r.sprintphase.sprint_id
}
func (r *sprintphaseResolver) PhaseId() int32 {
	return r.sprintphase.phase_id
}

// Mapper methods
func MapSprintPhase(modelSprintPhase models.SprintPhase) *sprintphase {

	if reflect.DeepEqual(modelSprintPhase, models.SprintPhase{}) {
		return &sprintphase{}
	}

	// Create graphql sprintphase from models SprintPhase
	sprintphase := sprintphase{
		id:        utils.UintToGraphId(modelSprintPhase.Id),
		phase_id:  int32(modelSprintPhase.PhaseId),
		sprint_id: int32(modelSprintPhase.SprintId),
	}
	return &sprintphase
}

// Reverse Mapper methods
func ReverseMapSprintPhase(mygraphqlSprintPhase *sprintphaseInput) models.SprintPhase {

	if reflect.DeepEqual(mygraphqlSprintPhase, sprintphaseInput{}) {
		return models.SprintPhase{}
	}

	// Create graphql sprintphase from models SprintPhase
	var sprintphaseModel models.SprintPhase
	if mygraphqlSprintPhase.Id == nil {
		sprintphaseModel = models.SprintPhase{}
	} else {
		sprintphaseModel = models.SprintPhase{
			Id:       utils.ConvertId(*mygraphqlSprintPhase.Id),
			PhaseId:  utils.Int32ToUint(*mygraphqlSprintPhase.PhaseId),
			SprintId: utils.Int32ToUint(*mygraphqlSprintPhase.SprintId),
		}
	}
	return sprintphaseModel
}
func ReverseMap2SprintPhase(structSprintPhase *sprintphase) models.SprintPhase {

	if reflect.DeepEqual(structSprintPhase, sprintphase{}) {
		return models.SprintPhase{}
	}

	// Create graphql sprintphase from models SprintPhase
	modelSprintPhase := models.SprintPhase{
		Id:       utils.ConvertId(structSprintPhase.id),
		PhaseId:  uint(structSprintPhase.phase_id),
		SprintId: uint(structSprintPhase.sprint_id),
	}
	return modelSprintPhase
}
