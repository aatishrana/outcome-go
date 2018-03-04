package mygraphql

import graphqlgo "github.com/neelance/graphql-go"

type Resolver struct{}

// query resolver for Org
func (r *Resolver) Org(args struct {
	ID graphqlgo.ID
}) []*orgResolver {
	return ResolveOrg(args)
}

// create resolver for Org
func (r *Resolver) UpsertOrg(args *struct {
	Org *orgInput
}) *orgResolver {
	return ResolveCreateOrg(args)
}

// delete resolver for Org
func (r *Resolver) DeleteOrg(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteOrg(args, "")
}

// query resolver for User
func (r *Resolver) User(args struct {
	ID graphqlgo.ID
}) []*userResolver {
	return ResolveUser(args)
}

// create resolver for User
func (r *Resolver) UpsertUser(args *struct {
	User *userInput
}) *userResolver {
	return ResolveCreateUser(args)
}

// delete resolver for User
func (r *Resolver) DeleteUser(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteUser(args, "")
}

// query resolver for Team
func (r *Resolver) Team(args struct {
	ID graphqlgo.ID
}) []*teamResolver {
	return ResolveTeam(args)
}

// create resolver for Team
func (r *Resolver) UpsertTeam(args *struct {
	Team *teamInput
}) *teamResolver {
	return ResolveCreateTeam(args)
}

// delete resolver for Team
func (r *Resolver) DeleteTeam(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteTeam(args, "")
}

// query resolver for Product
func (r *Resolver) Product(args struct {
	ID graphqlgo.ID
}) []*productResolver {
	return ResolveProduct(args)
}

// create resolver for Product
func (r *Resolver) UpsertProduct(args *struct {
	Product *productInput
}) *productResolver {
	return ResolveCreateProduct(args)
}

// delete resolver for Product
func (r *Resolver) DeleteProduct(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteProduct(args, "")
}

// query resolver for ProductBackLog
func (r *Resolver) ProductBackLog(args struct {
	ID graphqlgo.ID
}) []*productbacklogResolver {
	return ResolveProductBackLog(args)
}

// create resolver for ProductBackLog
func (r *Resolver) UpsertProductBackLog(args *struct {
	ProductBackLog *productbacklogInput
}) *productbacklogResolver {
	return ResolveCreateProductBackLog(args)
}

// delete resolver for ProductBackLog
func (r *Resolver) DeleteProductBackLog(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteProductBackLog(args, "")
}

// query resolver for Project
func (r *Resolver) Project(args struct {
	ID graphqlgo.ID
}) []*projectResolver {
	return ResolveProject(args)
}

// create resolver for Project
func (r *Resolver) UpsertProject(args *struct {
	Project *projectInput
}) *projectResolver {
	return ResolveCreateProject(args)
}

// delete resolver for Project
func (r *Resolver) DeleteProject(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteProject(args, "")
}

// query resolver for Story
func (r *Resolver) Story(args struct {
	ID graphqlgo.ID
}) []*storyResolver {
	return ResolveStory(args)
}

// create resolver for Story
func (r *Resolver) UpsertStory(args *struct {
	Story *storyInput
}) *storyResolver {
	return ResolveCreateStory(args)
}

// delete resolver for Story
func (r *Resolver) DeleteStory(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteStory(args, "")
}

// query resolver for Sprint
func (r *Resolver) Sprint(args struct {
	ID graphqlgo.ID
}) []*sprintResolver {
	return ResolveSprint(args)
}

// create resolver for Sprint
func (r *Resolver) UpsertSprint(args *struct {
	Sprint *sprintInput
}) *sprintResolver {
	return ResolveCreateSprint(args)
}

// delete resolver for Sprint
func (r *Resolver) DeleteSprint(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteSprint(args, "")
}

// query resolver for Phase
func (r *Resolver) Phase(args struct {
	ID graphqlgo.ID
}) []*phaseResolver {
	return ResolvePhase(args)
}

// create resolver for Phase
func (r *Resolver) UpsertPhase(args *struct {
	Phase *phaseInput
}) *phaseResolver {
	return ResolveCreatePhase(args)
}

// delete resolver for Phase
func (r *Resolver) DeletePhase(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeletePhase(args, "")
}

// query resolver for Task
func (r *Resolver) Task(args struct {
	ID graphqlgo.ID
}) []*taskResolver {
	return ResolveTask(args)
}

// create resolver for Task
func (r *Resolver) UpsertTask(args *struct {
	Task *taskInput
}) *taskResolver {
	return ResolveCreateTask(args)
}

// delete resolver for Task
func (r *Resolver) DeleteTask(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteTask(args, "")
}

// query resolver for UserTeam
func (r *Resolver) UserTeam(args struct {
	ID graphqlgo.ID
}) []*userteamResolver {
	return ResolveUserTeam(args)
}

// create resolver for UserTeam
func (r *Resolver) UpsertUserTeam(args *struct {
	UserTeam *userteamInput
}) *userteamResolver {
	return ResolveCreateUserTeam(args)
}

// delete resolver for UserTeam
func (r *Resolver) DeleteUserTeam(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteUserTeam(args, "")
}

// query resolver for SprintPhase
func (r *Resolver) SprintPhase(args struct {
	ID graphqlgo.ID
}) []*sprintphaseResolver {
	return ResolveSprintPhase(args)
}

// create resolver for SprintPhase
func (r *Resolver) UpsertSprintPhase(args *struct {
	SprintPhase *sprintphaseInput
}) *sprintphaseResolver {
	return ResolveCreateSprintPhase(args)
}

// delete resolver for SprintPhase
func (r *Resolver) DeleteSprintPhase(args struct {
	ID            graphqlgo.ID
	CascadeDelete bool
}) *int32 {
	return ResolveDeleteSprintPhase(args, "")
}
