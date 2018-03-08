package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to Team
func init() {

	// Standard routes
	router.Get("/team", GetAllTeams)
	router.Get("/team/:id", GetTeam)
	router.Post("/team", PostTeam)
	router.Put("/team/:id", PutTeam)
	router.Delete("/team/:id", DeleteTeam)

	router.Get("/team/:id/project", GetAllTeamProjects)
}
func GetAllTeams(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllTeams(GetLimitOffset(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetTeam(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTeam(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostTeam(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Team
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostTeam(data)
	json.NewEncoder(w).Encode(data)
}

func PutTeam(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Team
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutTeam(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteTeam(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteTeam(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}

func GetAllTeamProjects(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetProjectOfTeam(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}
