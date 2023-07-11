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

type RelationProgress struct {
	ChunkLen int
	CioRet   []byte
}

type RelationReturn struct {
	Progress   <-chan RelationProgress
	TotalItems int
}
