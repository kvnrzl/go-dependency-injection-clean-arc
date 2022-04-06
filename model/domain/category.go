package domain

import "time"

type Category struct {
	Id         int
	Name       string
	CreateTime time.Time
	UpdateTime time.Time
}
