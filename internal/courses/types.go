package courses

import (
	"cio_objects/internal/bqtools"
	"cio_objects/internal/paths"
)

type Course struct {
	ID                 string  `json:"id"`
	Slug               *string `json:"slug"`
	Name               *string `json:"name"`
	ShortName          *string `json:"short_name"`
	Type               *string `json:"type"`
	Status             *string `json:"status"`
	Active             *bool   `json:"active"`
	Published          *bool   `json:"published"`
	PublishedAt        *string `json:"published_at"`
	Created            *string `json:"created"`
	Modified           *string `json:"modified"`
	ExternalID         *string `json:"external_id"`
	CmsID              *string `json:"cms_id"`
	Duration           *int64  `json:"duration"`
	StartAt            *string `json:"start_at"`
	NotForSale         *bool   `json:"not_for_sale"`
	GraduationCriteria *string `json:"graduation_criteria"`
}

type UserWithActiveCourses struct {
	UserID     string   `json:"user_id"`
	Email      string   `json:"email"`
	CoursesIds []string `json:"courses"`
}

type CioCourses struct {
	region            paths.Region
	bq                *bqtools.BQ
	regionLowerText   string
	courseTable       string
	enrollmentTable   string
	userTable         string
	nomenclatureTable string
}
