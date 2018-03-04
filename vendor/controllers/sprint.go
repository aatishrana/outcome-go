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
}
func GetAllSprints(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllSprints()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetSprint(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetSprint(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PostSprint(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Sprint
	err := decoder.Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostSprint(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PutSprint(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Sprint
	err := decoder.Decode(&newData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutSprint(newData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DeleteSprint(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteSprint(utils.StringToUInt(ID), "")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
