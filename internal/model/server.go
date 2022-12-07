package model

type ServerListItem struct {
	ZoneId int    `json:"zoneId"`
	Url    string `json:"url"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
