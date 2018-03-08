package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to UserTeam
func init() {

	// Standard routes
	router.Get("/userteam", GetAllUserTeams)
	router.Get("/userteam/:id", GetUserTeam)
	router.Post("/userteam", PostUserTeam)
	router.Put("/userteam/:id", PutUserTeam)
	router.Delete("/userteam/:id", DeleteUserTeam)
}
func GetAllUserTeams(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllUserTeams(GetLimitOffset(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetUserTeam(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetUserTeam(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostUserTeam(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.UserTeam
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostUserTeam(data)
	json.NewEncoder(w).Encode(data)
}

func PutUserTeam(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.UserTeam
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutUserTeam(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteUserTeam(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteUserTeam(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}
