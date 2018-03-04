package mygraphql

import (
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type project struct {
	id         graphqlgo.ID
	name       string
	user_id    int32
	team_id    int32
	product_id int32
}

// Struct for upserting
type projectInput struct {
	Id        *graphqlgo.ID
	Name      string
	UserId    *int32
	TeamId    *int32
	ProductId *int32
}

// Struct for response
type projectResolver struct {
	project *project
}

func ResolveProject(args struct {
	ID graphqlgo.ID
}) (response []*projectResolver) {
	if args.ID != "" {
		response = append(response, &projectResolver{project: MapProject(models.GetProject(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllProjects() {
		response = append(response, &projectResolver{project: MapProject(val)})
	}
	return response
}

func ResolveCreateProject(args *struct {
	Project *projectInput
}) *projectResolver {
	var project = &project{}

	if args.Project.Id == nil {
		project = MapProject(models.PostProject(ReverseMapProject(args.Project)))
	} else {
		project = MapProject(models.PutProject(ReverseMapProject(args.Project)))
	}
	return &projectResolver{project}
}

// For Delete
func ResolveDeleteProject(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.ProjectChildren) == 0 && len(models.ProjectInterRelation) == 0 {
		del = models.DeleteProject(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	return response
}

// Fields resolvers
func (r *projectResolver) Id() graphqlgo.ID {
	return r.project.id
}
func (r *projectResolver) Name() string {
	return r.project.name
}
func (r *projectResolver) UserId() int32 {
	return r.project.user_id
}
func (r *projectResolver) TeamId() int32 {
	return r.project.team_id
}
func (r *projectResolver) ProductId() int32 {
	return r.project.product_id
}

// Mapper methods
func MapProject(modelProject models.Project) *project {

	if reflect.DeepEqual(modelProject, models.Project{}) {
		return &project{}
	}

	// Create graphql project from models Project
	project := project{
		id:         utils.UintToGraphId(modelProject.Id),
		name:       modelProject.Name,
		product_id: int32(modelProject.ProductId),
		team_id:    int32(modelProject.TeamId),
		user_id:    int32(modelProject.UserId),
	}
	return &project
}

// Reverse Mapper methods
func ReverseMapProject(mygraphqlProject *projectInput) models.Project {

	if reflect.DeepEqual(mygraphqlProject, projectInput{}) {
		return models.Project{}
	}

	// Create graphql project from models Project
	var projectModel models.Project
	if mygraphqlProject.Id == nil {
		projectModel = models.Project{Name: mygraphqlProject.Name}
	} else {
		projectModel = models.Project{
			Id:        utils.ConvertId(*mygraphqlProject.Id),
			Name:      mygraphqlProject.Name,
			ProductId: utils.Int32ToUint(*mygraphqlProject.ProductId),
			TeamId:    utils.Int32ToUint(*mygraphqlProject.TeamId),
			UserId:    utils.Int32ToUint(*mygraphqlProject.UserId),
		}
	}
	return projectModel
}
func ReverseMap2Project(structProject *project) models.Project {

	if reflect.DeepEqual(structProject, project{}) {
		return models.Project{}
	}

	// Create graphql project from models Project
	modelProject := models.Project{
		Id:        utils.ConvertId(structProject.id),
		Name:      structProject.name,
		ProductId: uint(structProject.product_id),
		TeamId:    uint(structProject.team_id),
		UserId:    uint(structProject.user_id),
	}
	return modelProject
}
