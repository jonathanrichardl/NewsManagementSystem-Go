package models

type NewsUpdate struct {
	Title  *string   `json:"title,omitempty"`
	Topic  *string   `json:"topic,omitempty"`
	Tags   *[]string `json:"tags,omitempty"`
	Status *string   `json:"status,omitempty"`
}
