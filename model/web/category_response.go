package web

import "time"

type CategoryResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
