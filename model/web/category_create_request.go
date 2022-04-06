package web

import "time"

type CategoryCreateRequest struct {
	Name       string    `validate:"required,max=200,min=1" json:"name"`
	CreateTime time.Time `validate:"required" json:"create_time"`
	UpdateTime time.Time `validate:"required" json:"update_time"`
}
