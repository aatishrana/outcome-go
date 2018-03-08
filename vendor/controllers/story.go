package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to Story
func init() {

	// Standard routes
	router.Get("/story", GetAllStorys)
	router.Get("/story/:id", GetStory)
	router.Post("/story", PostStory)
	router.Put("/story/:id", PutStory)
	router.Delete("/story/:id", DeleteStory)

	router.Get("/story/:id/tasks", GetStorysAllTasks)
}
func GetAllStorys(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllStorys(GetLimitOffset(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetStory(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetStory(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostStory(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Story
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostStory(data)
	json.NewEncoder(w).Encode(data)
}

func PutStory(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Story
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutStory(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteStory(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteStory(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}

func GetStorysAllTasks(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTasksOfStory(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}
