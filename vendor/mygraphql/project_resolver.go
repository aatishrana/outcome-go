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
type project struct {
	id         graphqlgo.ID
	name       string
	user_id    int32
	team_id    int32
	product_id int32
	storys     []*story
	sprints    []*sprint
	user       *user
	team       *team
	product    *product
}

// Struct for upserting
type projectInput struct {
	Id        *graphqlgo.ID
	Name      *string
	UserId    *int32
	TeamId    *int32
	ProductId *int32
	Storys    *[]storyInput
	Sprints   *[]sprintInput
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
	for _, val := range models.GetAllProjects(utils.GetDefaultLimitOffset()) {
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
	if project != nil && args.Project.Storys != nil {
		for _, dev := range *args.Project.Storys {
			if dev.Id == nil {
				story := ReverseMapStory(&dev)
				if story.ProjectId != 0 && utils.ConvertId(project.id) != story.ProjectId {
					// todo throw error
					return &projectResolver{}
				}
				story.ProjectId = utils.ConvertId(project.id)
				project.storys = append(project.storys, MapStory(models.PostStory(story)))
			} else {
				story := ReverseMapStory(&dev)
				if story.ProjectId != 0 && utils.ConvertId(project.id) != story.ProjectId {
					// todo throw error
					return &projectResolver{}
				}
				story.ProjectId = utils.ConvertId(project.id)
				project.storys = append(project.storys, MapStory(models.PutStory(story)))
			}
		}
	}
	if project != nil && args.Project.Sprints != nil {
		for _, dev := range *args.Project.Sprints {
			if dev.Id == nil {
				sprint := ReverseMapSprint(&dev)
				if sprint.ProjectId != 0 && utils.ConvertId(project.id) != sprint.ProjectId {
					// todo throw error
					return &projectResolver{}
				}
				sprint.ProjectId = utils.ConvertId(project.id)
				project.sprints = append(project.sprints, MapSprint(models.PostSprint(sprint)))
			} else {
				sprint := ReverseMapSprint(&dev)
				if sprint.ProjectId != 0 && utils.ConvertId(project.id) != sprint.ProjectId {
					// todo throw error
					return &projectResolver{}
				}
				sprint.ProjectId = utils.ConvertId(project.id)
				project.sprints = append(project.sprints, MapSprint(models.PutSprint(sprint)))
			}
		}
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
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.Project
		database.SQL.Model(models.Project{}).Preload("Storys").Preload("Sprints").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		for _, v := range data.Storys {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteStory(args, "")
			count++
		}

		for _, v := range data.Sprints {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteSprint(args, "")
			count++
		}

		del = models.DeleteProject(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.Project
	database.SQL.Model(models.Project{}).Preload("Storys").Preload("Sprints").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	for _, v := range data.Storys {
		if v.Id != 0 {
			flag++
		}
	}

	for _, v := range data.Sprints {
		if v.Id != 0 {
			flag++
		}
	}

	if flag == 0 {
		del = models.DeleteProject(utils.ConvertId(tempID), name)
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
func (r *projectResolver) Storys() []*storyResolver {
	var storys []*storyResolver
	if r.project != nil {
		story := models.GetStorysOfProject(utils.ConvertId(r.project.id))
		for _, value := range story {
			storys = append(storys, &storyResolver{MapStory(value)})
		}
		return storys
	}
	for _, value := range r.project.storys {
		storys = append(storys, &storyResolver{value})
	}
	return storys
}
func (r *projectResolver) Sprints() []*sprintResolver {
	var sprints []*sprintResolver
	if r.project != nil {
		sprint := models.GetSprintsOfProject(utils.ConvertId(r.project.id))
		for _, value := range sprint {
			sprints = append(sprints, &sprintResolver{MapSprint(value)})
		}
		return sprints
	}
	for _, value := range r.project.sprints {
		sprints = append(sprints, &sprintResolver{value})
	}
	return sprints
}
func (r *projectResolver) User() *userResolver {
	if r.project != nil {
		user := models.GetUserOfProject(ReverseMap2Project(r.project))
		return &userResolver{MapUser(user)}
	}
	return &userResolver{r.project.user}
}
func (r *projectResolver) Team() *teamResolver {
	if r.project != nil {
		team := models.GetTeamOfProject(ReverseMap2Project(r.project))
		return &teamResolver{MapTeam(team)}
	}
	return &teamResolver{r.project.team}
}
func (r *projectResolver) Product() *productResolver {
	if r.project != nil {
		product := models.GetProductOfProject(ReverseMap2Project(r.project))
		return &productResolver{MapProduct(product)}
	}
	return &productResolver{r.project.product}
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

	var id uint = 0
	var teamId uint = 0
	var userId uint = 0
	var productId uint = 0
	var name string = ""

	if mygraphqlProject.Id != nil {
		id = utils.ConvertId(*mygraphqlProject.Id)
	}
	if mygraphqlProject.TeamId != nil {
		teamId = utils.Int32ToUint(*mygraphqlProject.TeamId)
	}
	if mygraphqlProject.UserId != nil {
		userId = utils.Int32ToUint(*mygraphqlProject.UserId)
	}
	if mygraphqlProject.ProductId != nil {
		productId = utils.Int32ToUint(*mygraphqlProject.ProductId)
	}
	if mygraphqlProject.Name != nil {
		name = *mygraphqlProject.Name
	}

	return models.Project{
		Id:        id,
		Name:      name,
		ProductId: productId,
		TeamId:    teamId,
		UserId:    userId,
	}
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
