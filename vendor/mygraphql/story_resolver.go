package mygraphql

import (
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
