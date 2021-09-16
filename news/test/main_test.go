package test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"news/pkg/database"
	"news/pkg/handlers"
	"news/pkg/logger"
	"news/pkg/models"
	"news/pkg/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Init() *handlers.HTTPHandler {
	Logger := logger.NewLogger("log.txt")
	Logger.InfoLogger.Println("Initializing Program")
	Database, err := database.NewDatabase("mysql",
		"root", "123jonathan123100300!!!", "localhost:3306",
		"testers")

	if err != nil {
		log.Printf("Error received : %s\n", err.Error())
	}
	Router := router.NewRouterInstance()
	handlers := handlers.NewHttpHandlers(Database, Router, Logger)
	return handlers

}

func TestAddNewsHandler(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news", app.AddNewsHandler, http.MethodPost)
	newNews := models.News{
		Title:  "WhatisGo?",
		Tags:   []string{"Programming", "Go"},
		Topic:  "Technology",
		Status: "Draft",
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodPost, "/news", payload)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestRetreiveNews(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news/{title}", app.GetNewsByTitleHandler, http.MethodGet)
	req, _ := http.NewRequest(http.MethodGet, "/news/WhatIsGo", nil)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")
}

func TestDeleteNews(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news/{title}", app.DeleteNewsHandler, http.MethodDelete)
	req, _ := http.NewRequest(http.MethodDelete, "/news/WhatIsGo", nil)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestRetreiveNewsByTopic(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news/topic/{topic}", app.GetNewsByTopicHandler, http.MethodGet)
	req, _ := http.NewRequest(http.MethodGet, "/news/topic/business", nil)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestRetreiveNewsByState(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news/status/{status}", app.GetNewsByStatusHandler, http.MethodGet)
	req, _ := http.NewRequest(http.MethodGet, "/news/status/published", nil)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestModifyExistingNews(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news/{title}", app.ModifyNewsHandler, http.MethodPatch)
	Title := "NewTitle"
	Tags := []string{"Golang", "Program"}
	Topic := "Programming"
	newNews := models.NewsUpdate{
		Title: &Title,
		Tags:  &Tags,
		Topic: &Topic,
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodPatch, "/news/HealthyInvestation", payload)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestAddNewTagsIntoNews(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news/{title}/tags", app.AddNewTags, http.MethodPost)
	newNews := models.TagsUpdate{
		Tags: []string{"Golang", "Program"},
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodPost, "/news/WhatIsGo/tags", payload)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}

func TestDeleteTagsofNews(t *testing.T) {
	app := Init()
	app.Router.RegisterHandler("/news/{title}/tags/{tags}", app.AddNewTags, http.MethodDelete)
	newNews := models.TagsUpdate{
		Tags: []string{"Golang", "Program"},
	}
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(newNews)
	req, _ := http.NewRequest(http.MethodDelete, "/news/WhatIsGo/tags/Golang", payload)
	resp := httptest.NewRecorder()
	app.Router.Router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "Ok response")

}
