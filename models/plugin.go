package models

type Plugin struct {
	Model
	Name   string `json:"name"`
	Type   string `json:"type"`
	Config string `json:"config"`
}
