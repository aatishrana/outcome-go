package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to Task
func init() {

	// Standard routes
	router.Get("/task", GetAllTasks)
	router.Get("/task/:id", GetTask)
	router.Post("/task", PostTask)
	router.Put("/task/:id", PutTask)
	router.Delete("/task/:id", DeleteTask)
}
func GetAllTasks(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllTasks(GetLimitOffset(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetTask(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTask(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostTask(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Task
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostTask(data)
	json.NewEncoder(w).Encode(data)
}

func PutTask(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Task
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutTask(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteTask(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteTask(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}
