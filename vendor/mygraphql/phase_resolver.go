package mygraphql

import (
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type phase struct {
	id      graphqlgo.ID
	name    string
	sprints []*sprint
}

// Struct for upserting
type phaseInput struct {
	Id   *graphqlgo.ID
	Name string
}

// Struct for response
type phaseResolver struct {
	phase *phase
}

func ResolvePhase(args struct {
	ID graphqlgo.ID
}) (response []*phaseResolver) {
	if args.ID != "" {
		response = append(response, &phaseResolver{phase: MapPhase(models.GetPhase(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllPhases() {
		response = append(response, &phaseResolver{phase: MapPhase(val)})
	}
	return response
}

func ResolveCreatePhase(args *struct {
	Phase *phaseInput
}) *phaseResolver {
	var phase = &phase{}

	if args.Phase.Id == nil {
		phase = MapPhase(models.PostPhase(ReverseMapPhase(args.Phase)))
	} else {
		phase = MapPhase(models.PutPhase(ReverseMapPhase(args.Phase)))
	}
	return &phaseResolver{phase}
}

// For Delete
func ResolveDeletePhase(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.PhaseChildren) == 0 && len(models.PhaseInterRelation) == 0 {
		del = models.DeletePhase(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	return response
}

// Fields resolvers
func (r *phaseResolver) Id() graphqlgo.ID {
	return r.phase.id
}
func (r *phaseResolver) Name() string {
	return r.phase.name
}
func (r *phaseResolver) Sprints() []*sprintResolver {
	var sprints []*sprintResolver
	if r.phase != nil {
		sprint := models.GetSprintsOfPhase(utils.ConvertId(r.phase.id))
		for _, value := range sprint {
			sprints = append(sprints, &sprintResolver{MapSprint(value)})
		}
		return sprints
	}
	for _, value := range r.phase.sprints {
		sprints = append(sprints, &sprintResolver{value})
	}
	return sprints
}

// Mapper methods
func MapPhase(modelPhase models.Phase) *phase {

	if reflect.DeepEqual(modelPhase, models.Phase{}) {
		return &phase{}
	}

	// Create graphql phase from models Phase
	phase := phase{
		id:   utils.UintToGraphId(modelPhase.Id),
		name: modelPhase.Name,
	}
	return &phase
}

// Reverse Mapper methods
func ReverseMapPhase(mygraphqlPhase *phaseInput) models.Phase {

	if reflect.DeepEqual(mygraphqlPhase, phaseInput{}) {
		return models.Phase{}
	}

	// Create graphql phase from models Phase
	var phaseModel models.Phase
	if mygraphqlPhase.Id == nil {
		phaseModel = models.Phase{Name: mygraphqlPhase.Name}
	} else {
		phaseModel = models.Phase{
			Id:   utils.ConvertId(*mygraphqlPhase.Id),
			Name: mygraphqlPhase.Name,
		}
	}
	return phaseModel
}
func ReverseMap2Phase(structPhase *phase) models.Phase {

	if reflect.DeepEqual(structPhase, phase{}) {
		return models.Phase{}
	}

	// Create graphql phase from models Phase
	modelPhase := models.Phase{
		Id:   utils.ConvertId(structPhase.id),
		Name: structPhase.name,
	}
	return modelPhase
}
