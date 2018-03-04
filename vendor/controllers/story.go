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
}
func GetAllStorys(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllStorys()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetStory(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetStory(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PostStory(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Story
	err := decoder.Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostStory(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PutStory(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Story
	err := decoder.Decode(&newData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutStory(newData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DeleteStory(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteStory(utils.StringToUInt(ID), "")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
