package cio

import (
	"bytes"
	"cio_objects/internal/courses"
	"cio_objects/internal/helpers"
	"encoding/json"
)

func (cio *CIO) CreateOrUpdateClasses(classes []courses.Course) []byte {
	object_type_id := cio.getCoursesObjectTypeId()
	var batch []map[string]any
	for _, c := range classes {
		var attributes map[string]any
		if b, err := json.Marshal(c); err != nil {
			panic(err)
		} else {
			if err := json.Unmarshal(b, &attributes); err != nil {
				panic(err)
			}
		}
		body := map[string]any{
			"identifiers": map[string]string{
				"object_type_id": object_type_id,
				"object_id":      c.ID,
			},
			"type":              "object",
			"action":            "identify",
			"attributes":        attributes,
			"cio_relationships": []any{},
		}
		batch = append(batch, body)
	}
	body := map[string]any{
		"batch": batch,
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	return cio.trackRequest("/api/v2/batch", bytes.NewReader(bodyJson))
}

func (cio *CIO) RelateUsersWithCourseObjects(
	lmsUsers []courses.UserWithActiveCourses,
	cioUsers []helpers.Customer,
) RelationReturn {
	cache := map[string][]string{}
	var relatedCustomers []relateCustomerItem
	for i := 0; i < len(lmsUsers); i++ {
		cache[lmsUsers[i].Email] = lmsUsers[i].CoursesIds
	}
	for i := 0; i < len(cioUsers); i++ {
		if ids, ok := cache[cioUsers[i].Email]; ok {
			relatedCustomers = append(relatedCustomers, relateCustomerItem{
				CioID:      cioUsers[i].CioID,
				CoursesIds: ids,
			})
		}
	}
	cache = nil
	lmsUsers = nil
	cioUsers = nil

	batchSize := 500
	chunked := helpers.ChunkByMaxLen(relatedCustomers, batchSize)
	progress := make(chan RelationProgress)
	object_type_id := cio.getCoursesObjectTypeId()
	go func() {
		defer close(progress)
		for _, chunk := range chunked {
			var batch []map[string]any
			for _, c := range chunk {
				var cio_relationships []map[string]map[string]string
				for _, id := range c.CoursesIds {
					cio_relationships = append(cio_relationships, map[string]map[string]string{
						"identifiers": {
							"object_type_id": object_type_id,
							"object_id":      id,
						},
					})
				}
				body := map[string]any{
					"type": "person",
					"identifiers": map[string]string{
						"cio_id": c.CioID,
					},
					"action":            "identify",
					"cio_relationships": cio_relationships,
				}
				batch = append(batch, body)
			}
			body := map[string]any{
				"batch": batch,
			}
			bodyJson, err := json.Marshal(body)
			if err != nil {
				panic(err)
			}
			ret := cio.trackRequest("/api/v2/batch", bytes.NewReader(bodyJson))
			progress <- RelationProgress{
				ChunkLen: len(chunk),
				CioRet:   ret,
			}
		}
	}()
	return RelationReturn{
		TotalItems: len(relatedCustomers),
		Progress:   progress,
	}
}
