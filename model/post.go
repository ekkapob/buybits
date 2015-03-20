package model

import (
	"time"
)

type Post struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Skills      string    `json:"skills,omitempty"`
	Budget      string    `json:"budget,omitempty"`
	Owner       string    `json:"owner"`
}
