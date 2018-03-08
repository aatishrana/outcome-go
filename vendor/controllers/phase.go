package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to Phase
func init() {

	// Standard routes
	router.Get("/phase", GetAllPhases)
	router.Get("/phase/:id", GetPhase)
	router.Post("/phase", PostPhase)
	router.Put("/phase/:id", PutPhase)
	router.Delete("/phase/:id", DeletePhase)
}
func GetAllPhases(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllPhases()
	json.NewEncoder(w).Encode(data)
}

func GetPhase(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetPhase(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostPhase(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Phase
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostPhase(data)
	json.NewEncoder(w).Encode(data)
}

func PutPhase(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Phase
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutPhase(newData)
	json.NewEncoder(w).Encode(data)
}

func DeletePhase(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeletePhase(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}
