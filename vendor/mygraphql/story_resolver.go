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
type story struct {
	id                  graphqlgo.ID
	desc                string
	status              string
	point               int32
	product_back_log_id int32
	project_id          int32
	sprint_id           int32
	tasks               []*task
	productbacklog      *productbacklog
	project             *project
	sprint              *sprint
}

// Struct for upserting
type storyInput struct {
	Id               *graphqlgo.ID
	Desc             string
	Status           string
	Point            int32
	ProductBackLogId *int32
	ProjectId        *int32
	SprintId         *int32
	Tasks            *[]taskInput
}

// Struct for response
type storyResolver struct {
	story *story
}

func ResolveStory(args struct {
	ID graphqlgo.ID
}) (response []*storyResolver) {
	if args.ID != "" {
		response = append(response, &storyResolver{story: MapStory(models.GetStory(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllStorys() {
		response = append(response, &storyResolver{story: MapStory(val)})
	}
	return response
}

func ResolveCreateStory(args *struct {
	Story *storyInput
}) *storyResolver {
	var story = &story{}

	if args.Story.Id == nil {
		story = MapStory(models.PostStory(ReverseMapStory(args.Story)))
	} else {
		story = MapStory(models.PutStory(ReverseMapStory(args.Story)))
	}
	if story != nil && args.Story.Tasks != nil {
		for _, dev := range *args.Story.Tasks {
			if dev.Id == nil {
				task := ReverseMapTask(&dev)
				if task.StoryId != 0 && utils.ConvertId(story.id) != task.StoryId {
					// todo throw error
					return &storyResolver{}
				}
				task.StoryId = utils.ConvertId(story.id)
				story.tasks = append(story.tasks, MapTask(models.PostTask(task)))
			} else {
				task := ReverseMapTask(&dev)
				if task.StoryId != 0 && utils.ConvertId(story.id) != task.StoryId {
					// todo throw error
					return &storyResolver{}
				}
				task.StoryId = utils.ConvertId(story.id)
				story.tasks = append(story.tasks, MapTask(models.PutTask(task)))
			}
		}
	}
	return &storyResolver{story}
}

// For Delete
func ResolveDeleteStory(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.StoryChildren) == 0 && len(models.StoryInterRelation) == 0 {
		del = models.DeleteStory(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.Story
		database.SQL.Model(models.Story{}).Preload("Tasks").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		for _, v := range data.Tasks {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteTask(args, "")
			count++
		}

		del = models.DeleteStory(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.Story
	database.SQL.Model(models.Story{}).Preload("Tasks").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	for _, v := range data.Tasks {
		if v.Id != 0 {
			flag++
		}
	}

	if flag == 0 {
		del = models.DeleteStory(utils.ConvertId(tempID), name)
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
func (r *storyResolver) Id() graphqlgo.ID {
	return r.story.id
}
func (r *storyResolver) Desc() string {
	return r.story.desc
}
func (r *storyResolver) Status() string {
	return r.story.status
}
func (r *storyResolver) Point() int32 {
	return r.story.point
}
func (r *storyResolver) ProductBackLogId() int32 {
	return r.story.product_back_log_id
}
func (r *storyResolver) ProjectId() int32 {
	return r.story.project_id
}
func (r *storyResolver) SprintId() int32 {
	return r.story.sprint_id
}
func (r *storyResolver) Tasks() []*taskResolver {
	var tasks []*taskResolver
	if r.story != nil {
		task := models.GetTasksOfStory(utils.ConvertId(r.story.id))
		for _, value := range task {
			tasks = append(tasks, &taskResolver{MapTask(value)})
		}
		return tasks
	}
	for _, value := range r.story.tasks {
		tasks = append(tasks, &taskResolver{value})
	}
	return tasks
}
func (r *storyResolver) ProductBackLog() *productbacklogResolver {
	if r.story != nil {
		productbacklog := models.GetProductBackLogOfStory(ReverseMap2Story(r.story))
		return &productbacklogResolver{MapProductBackLog(productbacklog)}
	}
	return &productbacklogResolver{r.story.productbacklog}
}
func (r *storyResolver) Project() *projectResolver {
	if r.story != nil {
		project := models.GetProjectOfStory(ReverseMap2Story(r.story))
		return &projectResolver{MapProject(project)}
	}
	return &projectResolver{r.story.project}
}
func (r *storyResolver) Sprint() *sprintResolver {
	if r.story != nil {
		sprint := models.GetSprintOfStory(ReverseMap2Story(r.story))
		return &sprintResolver{MapSprint(sprint)}
	}
	return &sprintResolver{r.story.sprint}
}

// Mapper methods
func MapStory(modelStory models.Story) *story {

	if reflect.DeepEqual(modelStory, models.Story{}) {
		return &story{}
	}

	// Create graphql story from models Story
	story := story{
		desc:                modelStory.Desc,
		id:                  utils.UintToGraphId(modelStory.Id),
		point:               int32(modelStory.Point),
		product_back_log_id: int32(modelStory.ProductBackLogId),
		project_id:          int32(modelStory.ProjectId),
		sprint_id:           int32(modelStory.SprintId),
		status:              modelStory.Status,
	}
	return &story
}

// Reverse Mapper methods
func ReverseMapStory(mygraphqlStory *storyInput) models.Story {

	if reflect.DeepEqual(mygraphqlStory, storyInput{}) {
		return models.Story{}
	}

	// Create graphql story from models Story
	var storyModel models.Story
	if mygraphqlStory.Id == nil {
		storyModel = models.Story{
			Desc:   mygraphqlStory.Desc,
			Point:  utils.Int32ToUint(mygraphqlStory.Point),
			Status: mygraphqlStory.Status,
		}
	} else {
		storyModel = models.Story{
			Desc:             mygraphqlStory.Desc,
			Id:               utils.ConvertId(*mygraphqlStory.Id),
			Point:            utils.Int32ToUint(mygraphqlStory.Point),
			ProductBackLogId: utils.Int32ToUint(*mygraphqlStory.ProductBackLogId),
			ProjectId:        utils.Int32ToUint(*mygraphqlStory.ProjectId),
			SprintId:         utils.Int32ToUint(*mygraphqlStory.SprintId),
			Status:           mygraphqlStory.Status,
		}
	}
	return storyModel
}
func ReverseMap2Story(structStory *story) models.Story {

	if reflect.DeepEqual(structStory, story{}) {
		return models.Story{}
	}

	// Create graphql story from models Story
	modelStory := models.Story{
		Desc:             structStory.desc,
		Id:               utils.ConvertId(structStory.id),
		Point:            uint(structStory.point),
		ProductBackLogId: uint(structStory.product_back_log_id),
		ProjectId:        uint(structStory.project_id),
		SprintId:         uint(structStory.sprint_id),
		Status:           structStory.status,
	}
	return modelStory
}
