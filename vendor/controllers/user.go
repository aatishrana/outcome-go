package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to User
func init() {

	// Standard routes
	router.Get("/user", GetAllUsers)
	router.Get("/user/:id", GetUser)
	router.Post("/user", PostUser)
	router.Put("/user/:id", PutUser)
	router.Delete("/user/:id", DeleteUser)

	router.Get("/user/:id/team", GetUserCreatedTeam)
	router.Get("/user/:id/teams", GetUsersAllTeams)
	router.Get("/user/:id/product", GetUserCreatedProduct)
	router.Get("/user/:id/backlogs", GetUsersAllBacklogs)
	router.Get("/user/:id/project", GetUserCreatedProject)
	router.Get("/user/:id/tasks", GetUsersAllTasks)
}
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllUsers()
	json.NewEncoder(w).Encode(data)
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetUser(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.User
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostUser(data)
	json.NewEncoder(w).Encode(data)
}

func PutUser(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.User
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutUser(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteUser(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}

func GetUserCreatedTeam(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTeamOfUser(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func GetUsersAllTeams(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTeamsOfUser(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func GetUserCreatedProduct(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetProductOfUser(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func GetUsersAllBacklogs(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetProductBackLogsOfUser(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func GetUserCreatedProject(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetProjectOfUser(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func GetUsersAllTasks(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTasksOfUser(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}
