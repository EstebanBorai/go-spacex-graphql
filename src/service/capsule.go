package service

import (
	"log"
	"net/http"

	"github.com/estebanborai/go-spacex-graphql/src/util"

	"github.com/estebanborai/go-spacex-graphql/src/entity"
)

// CapsuleService provides methods to consume
// Capsule data
type CapsuleService struct {
	CapsulesURL string
	HTTPClient  *http.Client
}

// ProvideCapsuleService creates an instance of CapsuleService
func ProvideCapsuleService() *CapsuleService {
	client := new(http.Client)

	return &CapsuleService{
		CapsulesURL: "https://api.spacexdata.com/v3/capsules",
		HTTPClient:  client,
	}
}

// AllCapsules gathers all the registered capsules
func (s *CapsuleService) AllCapsules() ([]*entity.Capsule, error) {
	resp, err := s.HTTPClient.Get(s.CapsulesURL)

	if err != nil {
		log.Fatal(err)
	}

	var responseBody []*entity.Capsule

	responseBody = make([]*entity.Capsule, 0)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		util.JSONStruct(resp.Body, &responseBody)
	}

	return responseBody, nil
}
