package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to Sprint
func init() {

	// Standard routes
	router.Get("/sprint", GetAllSprints)
	router.Get("/sprint/:id", GetSprint)
	router.Post("/sprint", PostSprint)
	router.Put("/sprint/:id", PutSprint)
	router.Delete("/sprint/:id", DeleteSprint)

	router.Get("/sprint/:id/stories", GetSprinsAllStories)
	router.Get("/sprint/:id/tasks", GetSprintsAllTasks)
}
func GetAllSprints(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllSprints(GetLimitOffset(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetSprint(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetSprint(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostSprint(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Sprint
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostSprint(data)
	json.NewEncoder(w).Encode(data)
}

func PutSprint(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Sprint
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutSprint(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteSprint(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteSprint(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}

func GetSprinsAllStories(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetStorysOfSprint(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func GetSprintsAllTasks(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTasksOfSprint(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}
