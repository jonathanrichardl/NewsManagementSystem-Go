package models

type News struct {
	Title  string   `json:"title"`
	Topic  string   `json:"topic"`
	Tags   []string `json:"tags"`
	Status string   `json:"status"`
}
