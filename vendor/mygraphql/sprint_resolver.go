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
type sprint struct {
	id         graphqlgo.ID
	name       string
	start_dt   string
	end_dt     string
	project_id int32
	storys     []*story
	phases     []*phase
	tasks      []*task
	project    *project
}

// Struct for upserting
type sprintInput struct {
	Id        *graphqlgo.ID
	Name      string
	StartDt   string
	EndDt     string
	ProjectId *int32
	Storys    *[]storyInput
	Phases    *[]phaseInput
	Tasks     *[]taskInput
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
	for _, val := range models.GetAllSprints(utils.GetDefaultLimitOffset()) {
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
	if sprint != nil && args.Sprint.Storys != nil {
		for _, dev := range *args.Sprint.Storys {
			if dev.Id == nil {
				story := ReverseMapStory(&dev)
				if story.SprintId != 0 && utils.ConvertId(sprint.id) != story.SprintId {
					// todo throw error
					return &sprintResolver{}
				}
				story.SprintId = utils.ConvertId(sprint.id)
				sprint.storys = append(sprint.storys, MapStory(models.PostStory(story)))
			} else {
				story := ReverseMapStory(&dev)
				if story.SprintId != 0 && utils.ConvertId(sprint.id) != story.SprintId {
					// todo throw error
					return &sprintResolver{}
				}
				story.SprintId = utils.ConvertId(sprint.id)
				sprint.storys = append(sprint.storys, MapStory(models.PutStory(story)))
			}
		}
	}
	if sprint != nil && args.Sprint.Phases != nil {
		for _, dev := range *args.Sprint.Phases {
			if dev.Id == nil {
				phase := ReverseMapPhase(&dev)
				sprint.phases = append(sprint.phases, MapPhase(models.PostPhase(phase)))

				var data = sprintphaseInput{}
				sprintphase := ReverseMapSprintPhase(&data)
				sprintId := utils.ConvertId(sprint.id)
				var phaseId uint
				for _, val := range sprint.phases {
					phaseId = utils.ConvertId(val.id)
				}
				sprintphase.SprintId = sprintId
				sprintphase.PhaseId = phaseId
				models.PostSprintPhase(sprintphase)

			} else {
				phase := ReverseMapPhase(&dev)
				sprint.phases = append(sprint.phases, MapPhase(models.PutPhase(phase)))

				var data = sprintphaseInput{}
				sprintphase := ReverseMapSprintPhase(&data)
				sprintId := utils.ConvertId(sprint.id)
				var phaseId uint
				for _, val := range sprint.phases {
					phaseId = utils.ConvertId(val.id)
				}
				sprintphase.SprintId = sprintId
				sprintphase.PhaseId = phaseId
				models.PostSprintPhase(sprintphase)

			}
		}
	}
	if sprint != nil && args.Sprint.Tasks != nil {
		for _, dev := range *args.Sprint.Tasks {
			if dev.Id == nil {
				task := ReverseMapTask(&dev)
				if task.SprintId != 0 && utils.ConvertId(sprint.id) != task.SprintId {
					// todo throw error
					return &sprintResolver{}
				}
				task.SprintId = utils.ConvertId(sprint.id)
				sprint.tasks = append(sprint.tasks, MapTask(models.PostTask(task)))
			} else {
				task := ReverseMapTask(&dev)
				if task.SprintId != 0 && utils.ConvertId(sprint.id) != task.SprintId {
					// todo throw error
					return &sprintResolver{}
				}
				task.SprintId = utils.ConvertId(sprint.id)
				sprint.tasks = append(sprint.tasks, MapTask(models.PutTask(task)))
			}
		}
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
	tempID := args.ID
	if args.CascadeDelete == true {
		var data models.Sprint
		database.SQL.Model(models.Sprint{}).Preload("Storys").Preload("SprintPhases").Preload("Tasks").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
		for _, v := range data.Storys {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteStory(args, "")
			count++
		}

		for _, v := range data.SprintPhases {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteSprintPhase(args, "")
			count++
		}

		for _, v := range data.Tasks {
			args.ID = utils.UintToGraphId(v.Id)
			ResolveDeleteTask(args, "")
			count++
		}

		del = models.DeleteSprint(utils.ConvertId(tempID), name)
		count++
		response = &count
		return response
	}

	var flag int
	var data models.Sprint
	database.SQL.Model(models.Sprint{}).Preload("Storys").Preload("SprintPhases").Preload("Tasks").Where("id=?", utils.ConvertId(args.ID)).Find(&data)
	for _, v := range data.Storys {
		if v.Id != 0 {
			flag++
		}
	}

	for _, v := range data.SprintPhases {
		if v.Id != 0 {
			flag++
		}
	}

	for _, v := range data.Tasks {
		if v.Id != 0 {
			flag++
		}
	}

	if flag == 0 {
		del = models.DeleteSprint(utils.ConvertId(tempID), name)
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
func (r *sprintResolver) Storys() []*storyResolver {
	var storys []*storyResolver
	if r.sprint != nil {
		story := models.GetStorysOfSprint(utils.ConvertId(r.sprint.id))
		for _, value := range story {
			storys = append(storys, &storyResolver{MapStory(value)})
		}
		return storys
	}
	for _, value := range r.sprint.storys {
		storys = append(storys, &storyResolver{value})
	}
	return storys
}
func (r *sprintResolver) Phases() []*phaseResolver {
	var phases []*phaseResolver
	if r.sprint != nil {
		phase := models.GetPhasesOfSprint(utils.ConvertId(r.sprint.id))
		for _, value := range phase {
			phases = append(phases, &phaseResolver{MapPhase(value)})
		}
		return phases
	}
	for _, value := range r.sprint.phases {
		phases = append(phases, &phaseResolver{value})
	}
	return phases
}
func (r *sprintResolver) Tasks() []*taskResolver {
	var tasks []*taskResolver
	if r.sprint != nil {
		task := models.GetTasksOfSprint(utils.ConvertId(r.sprint.id))
		for _, value := range task {
			tasks = append(tasks, &taskResolver{MapTask(value)})
		}
		return tasks
	}
	for _, value := range r.sprint.tasks {
		tasks = append(tasks, &taskResolver{value})
	}
	return tasks
}
func (r *sprintResolver) Project() *projectResolver {
	if r.sprint != nil {
		project := models.GetProjectOfSprint(ReverseMap2Sprint(r.sprint))
		return &projectResolver{MapProject(project)}
	}
	return &projectResolver{r.sprint.project}
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
