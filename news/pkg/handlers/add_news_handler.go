package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"news/pkg/models"
)

func (h *HTTPHandler) AddNewsHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.InfoLogger.Println("Add news request received")
	var news models.News
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.ErrorLogger.Println("Error reading add news request: ", err.Error())

	}
	err = json.Unmarshal(req, &news)
	if err != nil {
		h.logger.ErrorLogger.Println("Error unmarshalling request data: ", err.Error())

	}
	if h.database.CheckIfExists(fmt.Sprintf("SELECT EXISTS(SELECT * FROM news WHERE title = '%s')", news.Title)) {
		h.logger.ErrorLogger.Printf("News with title %s already exists!\n", news.Title)
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(fmt.Sprintf("406 - News '%s' already exists", news.Title)))
		return
	}
	h.database.AddData(fmt.Sprintf("INSERT INTO news (Title,Topic,Status) VALUES ('%s','%s','%s');", news.Title, news.Topic, news.Status))
	h.addTags(news.Title, news.Tags)
	h.logger.InfoLogger.Printf("News '%s' is added into database\n", news.Title)
}
