package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"news/pkg/models"

	"github.com/gorilla/mux"
)

func (h *HTTPHandler) ModifyNewsHandler(w http.ResponseWriter, r *http.Request) {
	var newsUpdate models.NewsUpdate
	request := mux.Vars(r)
	newsTitle := request["title"]
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.ErrorLogger.Println("Error reading modify news request: ", err.Error())

	}
	err = json.Unmarshal(req, &newsUpdate)
	if err != nil {
		h.logger.ErrorLogger.Println("Error unmarshaling modify news request: ", err.Error())
	}
	if newsUpdate.Title != nil {

		h.modifyTitle(newsTitle, newsUpdate.Title)
		h.logger.InfoLogger.Printf("Updating title of %s into %s\n", newsTitle, *newsUpdate.Title)
		newsTitle = *newsUpdate.Title
	}
	if newsUpdate.Topic != nil {
		h.modifyTopic(newsTitle, newsUpdate.Topic)
		h.logger.InfoLogger.Printf("Updating topic of %s into %s\n", newsTitle, *newsUpdate.Topic)
	}
	if newsUpdate.Tags != nil {
		h.modifyTags(newsTitle, newsUpdate.Tags)
		h.logger.InfoLogger.Printf("Updating topic of %s into %s\n", newsTitle, *newsUpdate.Tags)
	}
	if newsUpdate.Status != nil {
		h.modifyStatus(newsTitle, newsUpdate.Status)
		h.logger.InfoLogger.Printf("Updating topic of %s into %s\n", newsTitle, *newsUpdate.Status)
	}

}

func (h *HTTPHandler) AddNewTags(w http.ResponseWriter, r *http.Request) {
	h.logger.InfoLogger.Println("Request to add new tags received")
	request := mux.Vars(r)
	newsTitle := request["title"]
	var newTags models.TagsUpdate
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.ErrorLogger.Println("Error reading modify tags request: ", err.Error())

	}
	err = json.Unmarshal(req, &newTags)
	if err != nil {
		h.logger.ErrorLogger.Println("Error unmarshaling modify tags request: ", err.Error())
	}
	var tagsToAdd []string
	for _, tag := range newTags.Tags {
		if h.database.CheckIfExists(fmt.Sprintf("SELECT EXISTS(SELECT * FROM tags WHERE tags = '%s' AND news_id = (SELECT id from news where title = '%s'))", tag, newsTitle)) {
			h.logger.WarningLogger.Printf("News %s already have tag %s", newsTitle, tag)
			continue
		}
		tagsToAdd = append(tagsToAdd, tag)
	}
	if len(tagsToAdd) == 0 {
		h.logger.WarningLogger.Println("No new tags added!")
		return
	}
	h.logger.InfoLogger.Println("New tags added!")
	h.addTags(newsTitle, tagsToAdd)

}

func (h *HTTPHandler) RemoveTags(w http.ResponseWriter, r *http.Request) {
	h.logger.InfoLogger.Println("Request to remove tags received")
	request := mux.Vars(r)
	newsTitle := request["title"]
	tagToRemove := request["tags"]
	if !h.database.CheckIfExists(fmt.Sprintf("SELECT EXISTS(SELECT * FROM tags WHERE tags = '%s' AND news_id = (SELECT id from news where title = '%s')", tagToRemove, newsTitle)) {
		h.logger.WarningLogger.Printf("News %s does not have tag %s", newsTitle, tagToRemove)
		return
	}
	h.logger.InfoLogger.Println("New tags added!")
	h.deleteTags(newsTitle, tagToRemove)

}
