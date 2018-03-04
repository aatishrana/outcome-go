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
}
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetUser(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PostUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.User
	err := decoder.Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostUser(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PutUser(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.User
	err := decoder.Decode(&newData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutUser(newData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteUser(utils.StringToUInt(ID), "")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
