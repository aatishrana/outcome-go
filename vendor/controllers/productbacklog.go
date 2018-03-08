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
	router.Get("/backlog", GetAllProductBackLogs)
	router.Get("/backlog/:id", GetProductBackLog)
	router.Post("/backlog", PostProductBackLog)
	router.Put("/backlog/:id", PutProductBackLog)
	router.Delete("/backlog/:id", DeleteProductBackLog)

	router.Get("/backlog/:id/stories", GetProductBackLogsAllStories)
}
func GetAllProductBackLogs(w http.ResponseWriter, req *http.Request) {
	data := models.GetAllProductBackLogs(GetLimitOffset(w, req))
	json.NewEncoder(w).Encode(data)
}

func GetProductBackLog(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetProductBackLog(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}

func PostProductBackLog(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data models.ProductBackLog
	err := decoder.Decode(&data)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()
	data = models.PostProductBackLog(data)
	json.NewEncoder(w).Encode(data)
}

func PutProductBackLog(w http.ResponseWriter, req *http.Request) {
	params := router.Params(req)
	ID := params.ByName("id")
	decoder := json.NewDecoder(req.Body)
	var newData models.ProductBackLog
	err := decoder.Decode(&newData)
	if err != nil {
		json.NewEncoder(w).Encode("invalid data")
		return
	}
	defer req.Body.Close()

	newData.Id = utils.StringToUInt(ID)
	data := models.PutProductBackLog(newData)
	json.NewEncoder(w).Encode(data)
}

func DeleteProductBackLog(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.DeleteProductBackLog(utils.StringToUInt(ID), "")
	json.NewEncoder(w).Encode(data)
}

func GetProductBackLogsAllStories(w http.ResponseWriter, req *http.Request) {
	// Get the parameter id
	params := router.Params(req)
	ID := params.ByName("id")
	data := models.GetStorysOfProductBackLog(utils.StringToUInt(ID))
	json.NewEncoder(w).Encode(data)
}
