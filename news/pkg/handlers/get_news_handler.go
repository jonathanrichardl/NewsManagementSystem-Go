package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news/pkg/models"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) GetAllNewsHandler(w http.ResponseWriter, r *http.Request) {
	retrievedData, err := h.database.RetrieveData("SELECT title,topic,status FROM news")
	if err != nil {
		h.logger.ErrorLogger.Println("Can't retrieve news: ", err.Error())
	}
	response := h.retrieveNews(retrievedData)
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) GetNewsByTopicHandler(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	topicName := request["topic"]
	retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT title,status,topic FROM NEWS WHERE TOPIC = '%s';", topicName))
	if err != nil {
		h.logger.ErrorLogger.Println("Can't retrieve news: ", err.Error())
	}
	response := h.retrieveNews(retrievedData)
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) GetNewsByTitleHandler(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	newsName := request["title"]
	retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT title,topic,status FROM news WHERE title = '%s';", newsName))
	if err != nil {
		h.logger.ErrorLogger.Println("Can't retrieve news: ", err.Error())
	}
	response := h.retrieveNews(retrievedData)
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandler) GetNewsByStatusHandler(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	status := request["status"]
	switch status {
	case "deleted":
		retrievedData, err := h.database.RetrieveData("SELECT * FROM deleted;")
		if err != nil {
			h.logger.ErrorLogger.Println("Can't retrieve deleted news: ", err.Error())
		}
		var response []models.DeletedNews
		for retrievedData.Data.Next() {
			var each models.DeletedNews
			err := retrievedData.Data.Scan(&each.Id, &each.Title)
			if err != nil {
				h.logger.ErrorLogger.Println("Can't retrieve data: ", err.Error())
			}
			response = append(response, each)
		}
		json.NewEncoder(w).Encode(response)

	default:
		retrievedData, err := h.database.RetrieveData(fmt.Sprintf("SELECT title,topic,status FROM news WHERE status = '%s';", status))
		if err != nil {
			h.logger.ErrorLogger.Println("Can't retrieve data: ", err.Error())
		}
		response := h.retrieveNews(retrievedData)
		json.NewEncoder(w).Encode(response)
	}

}
