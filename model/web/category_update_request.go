package web

import "time"

type CategoryUpdateRequest struct {
	Id         int       `validate:"required" json:"id"`
	Name       string    `validate:"required,max=200,min=1" json:"name"`
	UpdateTime time.Time `validate:"required" json:"update_time"`
}
