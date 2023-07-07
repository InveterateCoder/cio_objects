package cio

import "cio_objects/internal/paths"

type CIO struct {
	region      paths.Region
	trackURL    string
	user        string
	password    string
	bearerToken string
}

type relateCustomerItem struct {
	CioID      string   `json:"cio_id"`
	CoursesIds []string `json:"courses"`
}
