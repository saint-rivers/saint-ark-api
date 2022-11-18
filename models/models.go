package model

import "time"

type Resource struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	UploadDate time.Time `json:"uploadDate" example:"0000-00-00"`
	Format     string    `json:"format"`
}

type App struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"createdDate"`
	Owner       string    `json:"ownerId"`
	IsPublic    bool      `json:"isPublic"`
}
