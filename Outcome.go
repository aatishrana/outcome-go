package main

import (
	config "config"
	controllers "controllers"
	database "database"
	//graphqlgo "github.com/neelance/graphql-go"
	jsonconfig "jsonconfig"
	models "models"
	//mygraphql "mygraphql"
	os "os"
	route "route"
	runtime "runtime"
	server "server"
)

var conf = &config.Configuration{}

func init() {
	//  Use all cpu cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", conf)

	// Connect to database
	database.Connect(conf.Database)

	// Create schema
	//schema := graphqlgo.MustParseSchema(mygraphql.Schema, &mygraphql.Resolver{})

	// Load the controller routes
	controllers.Load(nil)

	// Auto migrate all models
	database.SQL.AutoMigrate(&models.Org{}, &models.User{}, &models.Team{}, &models.Product{}, &models.ProductBackLog{}, &models.Project{}, &models.Story{}, &models.Sprint{}, &models.Phase{}, &models.Task{}, &models.UserTeam{}, &models.SprintPhase{})

	// Start the listener
	server.Run(route.LoadHTTP(), route.LoadHTTPS(), conf.Server)
}
