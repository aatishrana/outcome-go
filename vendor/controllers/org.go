package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to Org
func init() {

	// Standard routes
	router.Get("/org", GetAllOrgs)
	router.Get("/org/:id", GetOrg)
	router.Post("/org", PostOrg)
	router.Put("/org/:id", PutOrg)
	router.Delete("/org/:id", DeleteOrg)

	router.Get("/org/:id/teams", GetAllOrgTeams)
	router.Get("/org/:id/users", GetAllOrgUsers)
	router.Get("/org/:id/products", GetAllOrgProducts)
}
func GetAllOrgs(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllOrgs()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetOrg(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetOrg(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PostOrg(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Org
	err := decoder.Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostOrg(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PutOrg(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.Org
	err := decoder.Decode(&newData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutOrg(newData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DeleteOrg(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteOrg(utils.StringToUInt(ID), "")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetAllOrgTeams(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetTeamsOfOrg(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetAllOrgUsers(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetUsersOfOrg(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetAllOrgProducts(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetProductsOfOrg(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
