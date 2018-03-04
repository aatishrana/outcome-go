package controllers

import (
	json "encoding/json"
	models "models"
	http "net/http"
	router "router"
	utils "utils"
)

// Routes related to ProductBackLog
func init() {

	// Standard routes
	router.Get("/productbacklog", GetAllProductBackLogs)
	router.Get("/productbacklog/:id", GetProductBackLog)
	router.Post("/productbacklog", PostProductBackLog)
	router.Put("/productbacklog/:id", PutProductBackLog)
	router.Delete("/productbacklog/:id", DeleteProductBackLog)
}
func GetAllProductBackLogs(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllProductBackLogs()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetProductBackLog(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetProductBackLog(utils.StringToUInt(ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PostProductBackLog(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.ProductBackLog
	err := decoder.Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostProductBackLog(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func PutProductBackLog(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.ProductBackLog
	err := decoder.Decode(&newData)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutProductBackLog(newData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DeleteProductBackLog(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteProductBackLog(utils.StringToUInt(ID), "")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
