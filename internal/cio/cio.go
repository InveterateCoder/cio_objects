package cio

import (
	awshelp "cio_objects/internal/aws_help"
	"cio_objects/internal/paths"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewCIO(region paths.Region) *CIO {
	secret := awshelp.GetSecret(fmt.Sprintf(
		"prod/customer-io/%s",
		strings.ToLower(string(region))))
	cio := CIO{
		region:      region,
		trackURL:    "https://track-eu.customer.io",
		user:        secret["trackingSiteId"],
		password:    secret["trackingApiKey"],
		bearerToken: secret["appBearerToken"],
	}
	return &cio
}

func (cio *CIO) getCoursesObjectTypeId() string {
	if cio.region == paths.BR {
		return "2"
	} else {
		return "1"
	}
}

func (cio *CIO) trackRequest(endpoint string, body io.Reader) []byte {
	req, err := http.NewRequest(
		http.MethodPost,
		cio.trackURL+endpoint,
		body,
	)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(cio.user, cio.password)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	ret, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return ret
}
