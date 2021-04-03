package models

type Comment struct {
	Model
	UserID  string `json:"user_id"`
	Comment string `json:"comment"`
}
