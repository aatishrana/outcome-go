package mygraphql

var Schema = ` 
 schema {
 	query: Query
 	mutation: Mutation
 }

 # The query type, represents all of the entry points into our object graph
 type Query {
 	org(id: ID!) : [Org]!
 	user(id: ID!) : [User]!
 	team(id: ID!) : [Team]!
 	product(id: ID!) : [Product]!
 	productbacklog(id: ID!) : [ProductBackLog]!
 	project(id: ID!) : [Project]!
 	story(id: ID!) : [Story]!
 	sprint(id: ID!) : [Sprint]!
 	phase(id: ID!) : [Phase]!
 	task(id: ID!) : [Task]!
 	userteam(id: ID!) : [UserTeam]!
 	sprintphase(id: ID!) : [SprintPhase]!
 }

 # The mutation type, represents all updates we can make to our data
 type Mutation {
 # Create
 	upsertOrg(org: OrgInput!) :Org
 	upsertUser(user: UserInput!) :User
 	upsertTeam(team: TeamInput!) :Team
 	upsertProduct(product: ProductInput!) :Product
 	upsertProductBackLog(productbacklog: ProductBackLogInput!) :ProductBackLog
 	upsertProject(project: ProjectInput!) :Project
 	upsertStory(story: StoryInput!) :Story
 	upsertSprint(sprint: SprintInput!) :Sprint
 	upsertPhase(phase: PhaseInput!) :Phase
 	upsertTask(task: TaskInput!) :Task
 	upsertUserTeam(userteam: UserTeamInput!) :UserTeam
 	upsertSprintPhase(sprintphase: SprintPhaseInput!) :SprintPhase
 # Delete
 	deleteOrg(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteUser(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteTeam(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteProduct(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteProductBackLog(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteProject(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteStory(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteSprint(id: ID!,cascadeDelete: Boolean!) : Int
 	deletePhase(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteTask(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteUserTeam(id: ID!,cascadeDelete: Boolean!) : Int
 	deleteSprintPhase(id: ID!,cascadeDelete: Boolean!) : Int
 }

 type Org {
 	id: ID!
 	name: String!
 	users:[User!]!
 	teams:[Team!]!
 	products:[Product!]!
 }
 input OrgInput {
 	id: ID 
 	name: String
 	users: [UserInput!]
 	teams: [TeamInput!]
 	products: [ProductInput!]
 }

 type User {
 	id: ID!
 	first_name: String!
 	last_name: String!
 	email: String!
 	password: String!
 	token: String!
 	org_id: Int!
 	teams:[Team!]!
 	team: Team!
 	product: Product!
 	productbacklogs:[ProductBackLog!]!
 	project: Project!
 	tasks:[Task!]!
 	org: Org!
 }
 input UserInput {
 	id: ID 
 	first_name: String
 	last_name: String
 	email: String
 	password: String
 	token: String
 	org_id: Int 
 	teams: [TeamInput!]
 	team: TeamInput
 	product: ProductInput
 	productbacklogs: [ProductBackLogInput!]
 	project: ProjectInput
 	tasks: [TaskInput!]
 }

 type Team {
 	id: ID!
 	name: String!
 	user_id: Int!
 	org_id: Int!
 	project: Project!
 	org: Org!
 	users: [User!]!
 	user: User!
 }
 input TeamInput {
 	id: ID 
 	name: String
 	user_id: Int 
 	org_id: Int 
 	project: ProjectInput
 }

 type Product {
 	id: ID!
 	name: String!
 	desc: String!
 	user_id: Int!
 	org_id: Int!
 	productbacklogs:[ProductBackLog!]!
 	projects:[Project!]!
 	org: Org!
 	user: User!
 }
 input ProductInput {
 	id: ID 
 	name: String
 	desc: String
 	user_id: Int 
 	org_id: Int 
 	productbacklogs: [ProductBackLogInput!]
 	projects: [ProjectInput!]
 }

 type ProductBackLog {
 	id: ID!
 	desc: String!
 	type_cd: String!
 	priority: String!
 	user_id: Int!
 	product_id: Int!
 	storys:[Story!]!
 	user: User!
 	product: Product!
 }
 input ProductBackLogInput {
 	id: ID 
 	desc: String
 	type_cd: String
 	priority: String
 	user_id: Int 
 	product_id: Int 
 	storys: [StoryInput!]
 }

 type Project {
 	id: ID!
 	name: String!
 	user_id: Int!
 	team_id: Int!
 	product_id: Int!
 	storys:[Story!]!
 	sprints:[Sprint!]!
 	user: User!
 	team: Team!
 	product: Product!
 }
 input ProjectInput {
 	id: ID 
 	name: String
 	user_id: Int 
 	team_id: Int 
 	product_id: Int 
 	storys: [StoryInput!]
 	sprints: [SprintInput!]
 }

 type Story {
 	id: ID!
 	desc: String!
 	status: String!
 	point: Int!
 	product_back_log_id: Int!
 	project_id: Int!
 	sprint_id: Int!
 	tasks:[Task!]!
 	productbacklog: ProductBackLog!
 	project: Project!
 	sprint: Sprint!
 }
 input StoryInput {
 	id: ID 
 	desc: String!
 	status: String!
 	point: Int!
 	product_back_log_id: Int 
 	project_id: Int 
 	sprint_id: Int 
 	tasks: [TaskInput!]
 }

 type Sprint {
 	id: ID!
 	name: String!
 	start_dt: String!
 	end_dt: String!
 	project_id: Int!
 	storys:[Story!]!
 	phases:[Phase!]!
 	tasks:[Task!]!
 	project: Project!
 }
 input SprintInput {
 	id: ID 
 	name: String
 	start_dt: String
 	end_dt: String
 	project_id: Int 
 	storys: [StoryInput!]
 	phases: [PhaseInput!]
 	tasks: [TaskInput!]
 }

 type Phase {
 	id: ID!
 	name: String!
 	sprints: [Sprint!]!
 }
 input PhaseInput {
 	id: ID 
 	name: String
 }

 type Task {
 	id: ID!
 	sprint_id: Int!
 	story_id: Int!
 	sprint_phase_id: Int!
 	user_id: Int!
 	point: Int!
 	start_dt_tm: String!
 	end_dt_tm: String!
 	user: User!
 	story: Story!
 	sprint: Sprint!
 }
 input TaskInput {
 	id: ID 
 	sprint_id: Int 
 	story_id: Int 
 	sprint_phase_id: Int 
 	user_id: Int 
 	point: Int!
 	start_dt_tm: String!
 	end_dt_tm: String!
 }

 type UserTeam {
 	id: ID!
 	user_id: Int!
 	team_id: Int!
 }
 input UserTeamInput {
 	id: ID 
 	user_id: Int 
 	team_id: Int 
 }

 type SprintPhase {
 	id: ID!
 	sprint_id: Int!
 	phase_id: Int!
 }
 input SprintPhaseInput {
 	id: ID 
 	sprint_id: Int 
 	phase_id: Int 
 }

`
