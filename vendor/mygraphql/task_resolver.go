package mygraphql

import (
	graphqlgo "github.com/neelance/graphql-go"
	models "models"
	reflect "reflect"
	utils "utils"
)

// Struct for graphql
type task struct {
	id              graphqlgo.ID
	sprint_id       int32
	story_id        int32
	sprint_phase_id int32
	user_id         int32
	point           int32
	start_dt_tm     string
	end_dt_tm       string
	user            *user
	story           *story
	sprint          *sprint
}

// Struct for upserting
type taskInput struct {
	Id            *graphqlgo.ID
	SprintId      *int32
	StoryId       *int32
	SprintPhaseId *int32
	UserId        *int32
	Point         int32
	StartDtTm     string
	EndDtTm       string
}

// Struct for response
type taskResolver struct {
	task *task
}

func ResolveTask(args struct {
	ID graphqlgo.ID
}) (response []*taskResolver) {
	if args.ID != "" {
		response = append(response, &taskResolver{task: MapTask(models.GetTask(utils.ConvertId(args.ID)))})
		return response
	}
	for _, val := range models.GetAllTasks(utils.GetDefaultLimitOffset()) {
		response = append(response, &taskResolver{task: MapTask(val)})
	}
	return response
}

func ResolveCreateTask(args *struct {
	Task *taskInput
}) *taskResolver {
	var task = &task{}

	if args.Task.Id == nil {
		task = MapTask(models.PostTask(ReverseMapTask(args.Task)))
	} else {
		task = MapTask(models.PutTask(ReverseMapTask(args.Task)))
	}
	return &taskResolver{task}
}

// For Delete
func ResolveDeleteTask(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}, name string) (response *int32) {
	var del bool
	var count int32
	if len(models.TaskChildren) == 0 && len(models.TaskInterRelation) == 0 {
		del = models.DeleteTask(utils.ConvertId(args.ID), name)
		if del == true {
			count++
		}
		response = &count
		return response
	}
	return response
}

// Fields resolvers
func (r *taskResolver) Id() graphqlgo.ID {
	return r.task.id
}
func (r *taskResolver) SprintId() int32 {
	return r.task.sprint_id
}
func (r *taskResolver) StoryId() int32 {
	return r.task.story_id
}
func (r *taskResolver) SprintPhaseId() int32 {
	return r.task.sprint_phase_id
}
func (r *taskResolver) UserId() int32 {
	return r.task.user_id
}
func (r *taskResolver) Point() int32 {
	return r.task.point
}
func (r *taskResolver) StartDtTm() string {
	return r.task.start_dt_tm
}
func (r *taskResolver) EndDtTm() string {
	return r.task.end_dt_tm
}
func (r *taskResolver) User() *userResolver {
	if r.task != nil {
		user := models.GetUserOfTask(ReverseMap2Task(r.task))
		return &userResolver{MapUser(user)}
	}
	return &userResolver{r.task.user}
}
func (r *taskResolver) Story() *storyResolver {
	if r.task != nil {
		story := models.GetStoryOfTask(ReverseMap2Task(r.task))
		return &storyResolver{MapStory(story)}
	}
	return &storyResolver{r.task.story}
}
func (r *taskResolver) Sprint() *sprintResolver {
	if r.task != nil {
		sprint := models.GetSprintOfTask(ReverseMap2Task(r.task))
		return &sprintResolver{MapSprint(sprint)}
	}
	return &sprintResolver{r.task.sprint}
}

// Mapper methods
func MapTask(modelTask models.Task) *task {

	if reflect.DeepEqual(modelTask, models.Task{}) {
		return &task{}
	}

	// Create graphql task from models Task
	task := task{
		end_dt_tm:       modelTask.EndDtTm,
		id:              utils.UintToGraphId(modelTask.Id),
		point:           int32(modelTask.Point),
		sprint_id:       int32(modelTask.SprintId),
		sprint_phase_id: int32(modelTask.SprintPhaseId),
		start_dt_tm:     modelTask.StartDtTm,
		story_id:        int32(modelTask.StoryId),
		user_id:         int32(modelTask.UserId),
	}
	return &task
}

// Reverse Mapper methods
func ReverseMapTask(mygraphqlTask *taskInput) models.Task {

	if reflect.DeepEqual(mygraphqlTask, taskInput{}) {
		return models.Task{}
	}

	// Create graphql task from models Task
	var taskModel models.Task
	if mygraphqlTask.Id == nil {
		taskModel = models.Task{
			EndDtTm:   mygraphqlTask.EndDtTm,
			Point:     utils.Int32ToUint(mygraphqlTask.Point),
			StartDtTm: mygraphqlTask.StartDtTm,
		}
	} else {
		taskModel = models.Task{
			EndDtTm:       mygraphqlTask.EndDtTm,
			Id:            utils.ConvertId(*mygraphqlTask.Id),
			Point:         utils.Int32ToUint(mygraphqlTask.Point),
			SprintId:      utils.Int32ToUint(*mygraphqlTask.SprintId),
			SprintPhaseId: utils.Int32ToUint(*mygraphqlTask.SprintPhaseId),
			StartDtTm:     mygraphqlTask.StartDtTm,
			StoryId:       utils.Int32ToUint(*mygraphqlTask.StoryId),
			UserId:        utils.Int32ToUint(*mygraphqlTask.UserId),
		}
	}
	return taskModel
}
func ReverseMap2Task(structTask *task) models.Task {

	if reflect.DeepEqual(structTask, task{}) {
		return models.Task{}
	}

	// Create graphql task from models Task
	modelTask := models.Task{
		EndDtTm:       structTask.end_dt_tm,
		Id:            utils.ConvertId(structTask.id),
		Point:         uint(structTask.point),
		SprintId:      uint(structTask.sprint_id),
		SprintPhaseId: uint(structTask.sprint_phase_id),
		StartDtTm:     structTask.start_dt_tm,
		StoryId:       uint(structTask.story_id),
		UserId:        uint(structTask.user_id),
	}
	return modelTask
}
