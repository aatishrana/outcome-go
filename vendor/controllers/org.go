package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
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
	data := models.GetAllOrgs(GetLimitOffset(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetOrg(w http.ResponseWriter, req *http.Request) {
	data := models.GetOrg(GetId(w, req))
	json.NewEncoder(w).Encode(data)
}

func PostOrg(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.Org
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostOrg(data)
	json.NewEncoder(w).Encode(data)
}

func PutOrg(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var newData models.Org
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = GetId(w, req)
	data := models.PutOrg(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteOrg(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	data := models.DeleteOrg(GetId(w, req), "")
	json.NewEncoder(w).Encode(data)
}

func GetAllOrgTeams(w http.ResponseWriter, req *http.Request) {
	data := models.GetTeamsOfOrg(GetId(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetAllOrgUsers(w http.ResponseWriter, req *http.Request) {
	data := models.GetUsersOfOrg(GetId(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetAllOrgProducts(w http.ResponseWriter, req *http.Request) {
	data := models.GetProductsOfOrg(GetId(w, req))
	json.NewEncoder(w).Encode(data)
}
