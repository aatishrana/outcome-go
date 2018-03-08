package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to SprintPhase
func init() {

	// Standard routes
	router.Get("/sprintphase", GetAllSprintPhases)
	router.Get("/sprintphase/:id", GetSprintPhase)
	router.Post("/sprintphase", PostSprintPhase)
	router.Put("/sprintphase/:id", PutSprintPhase)
	router.Delete("/sprintphase/:id", DeleteSprintPhase)
}
func GetAllSprintPhases(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllSprintPhases()
	json.NewEncoder(w).Encode(data)
}

func GetSprintPhase(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetSprintPhase(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostSprintPhase(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.SprintPhase
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostSprintPhase(data)
	json.NewEncoder(w).Encode(data)
}

func PutSprintPhase(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.SprintPhase
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutSprintPhase(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteSprintPhase(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteSprintPhase(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}
