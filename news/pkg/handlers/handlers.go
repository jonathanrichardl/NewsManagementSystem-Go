package handlers

import (
	"news/pkg/database"
	"news/pkg/logger"
	"news/pkg/router"
)

type HTTPHandler struct {
	database *database.DBInstance
	Router   *router.RouterInstance
	logger   *logger.LoggerInstance
}

func NewHttpHandlers(DatabaseInstance *database.DBInstance, RouterInstance *router.RouterInstance, LoggerInstance *logger.LoggerInstance) *HTTPHandler {
	return &HTTPHandler{
		database: DatabaseInstance,
		Router:   RouterInstance,
		logger:   LoggerInstance,
	}
}

func (h *HTTPHandler) RegisterAllHandlers() {
	h.Router.RegisterHandler("/news", h.AddNewsHandler, "POST")
	h.Router.RegisterHandler("/news", h.GetAllNewsHandler, "GET")
	h.Router.RegisterHandler("/news/{title}", h.GetNewsByTitleHandler, "GET")
	h.Router.RegisterHandler("/news/{title}", h.DeleteNewsHandler, "DELETE")
	h.Router.RegisterHandler("/news/{title}", h.ModifyNewsHandler, "PATCH")
	h.Router.RegisterHandler("/news/{title}/tags", h.AddNewTags, "POST")
	h.Router.RegisterHandler("/news/{title}/tags/{tags}", h.RemoveTags, "DELETE")
	h.Router.RegisterHandler("/news/topic/{topic}", h.GetNewsByTopicHandler, "GET")
	h.Router.RegisterHandler("/news/status/{status}", h.GetNewsByStatusHandler, "GET")
}
