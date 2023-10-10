package model

type ServerListItem struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Hot    int    `json:"hot"`
}
