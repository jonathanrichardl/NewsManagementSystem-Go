package handlers

import (
	"fmt"
	"log"
	"news/pkg/database"
	"news/pkg/models"
)

/*
For converting SQL raw data into Go types
*/
func (h *HTTPHandler) getTagsofNews(news *models.News) {
	retrievedTags, err := h.database.RetrieveData(fmt.Sprintf("SELECT tags FROM tags WHERE news_id = (SELECT id FROM news WHERE title = '%s');", news.Title))
	if err != nil {
		log.Println("Error retrieving tags : ", err.Error())
	}
	for retrievedTags.Data.Next() {
		var tags string
		err = retrievedTags.Data.Scan(&tags)
		if err != nil {
			log.Println("Error scanning tags: ", err.Error())
		}
		news.Tags = append(news.Tags, tags)
	}

}

func (h *HTTPHandler) retrieveNews(retrievedData *database.RetrievedData) []models.News {
	var response []models.News
	for retrievedData.Data.Next() {
		var each models.News
		err := retrievedData.Data.Scan(&each.Title, &each.Topic, &each.Status)
		if err != nil {
			log.Println("Error retrieving sql  ", err.Error())
		}

		h.getTagsofNews(&each)
		response = append(response, each)
	}
	return response
}

/*
For Modifying purposes
*/

func (h *HTTPHandler) modifyTopic(newsTitle string, newTopic *string) {
	err := h.database.UpdateData(fmt.Sprintf("UPDATE news SET topic = '%s' WHERE title = '%s';", *newTopic, newsTitle))
	if err != nil {
		log.Println("Error updating title : ", err.Error())
	}
}

func (h *HTTPHandler) modifyTags(newsTitle string, newTags *[]string) {
	err := h.database.DeleteData(fmt.Sprintf("DELETE FROM tags WHERE news_id = (SELECT id FROM news WHERE title = '%s');", newsTitle))
	if err != nil {
		log.Println("Error removing tags : ", err.Error())
	}
	h.addTags(newsTitle, *newTags)

}

func (h *HTTPHandler) modifyTitle(previousTitle string, newTitle *string) {
	err := h.database.UpdateData(fmt.Sprintf("UPDATE news SET title = '%s' WHERE title = '%s';", *newTitle, previousTitle))
	if err != nil {
		log.Println("Error updating title : ", err.Error())
	}
}

func (h *HTTPHandler) modifyStatus(newsTitle string, newStatus *string) {
	err := h.database.UpdateData(fmt.Sprintf("UPDATE news SET status = '%s' WHERE title = '%s';", *newStatus, newsTitle))
	if err != nil {
		log.Println("Error updating title : ", err.Error())
	}
}

func (h *HTTPHandler) addTags(newsTitle string, tags []string) {
	for _, tag := range tags {
		err := h.database.AddData(fmt.Sprintf("INSERT INTO tags (Tags,news_id) VALUES ('%s', (SELECT id from news WHERE title = '%s')); ", tag, newsTitle))
		if err != nil {
			log.Println("Error adding tags: ", err.Error())
		}
	}
}

func (h *HTTPHandler) deleteTags(newsTitle string, tags string) {
	err := h.database.DeleteData(fmt.Sprintf("DELETE FROM tags WHERE tags = '%s'AND news_id = (SELECT id FROM news WHERE title = '%s');", tags, newsTitle))
	if err != nil {
		log.Println("Error removing tags: ", err.Error())
	}
}
