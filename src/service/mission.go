package service

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/estebanborai/go-spacex-graphql/src/entity"
)

// MissionService provides methods to consume
// Mission data
type MissionService struct {
	MissionsURL string
	HTTPClient  *http.Client
}

// ProvideMissionService creates an instance of MissionService
func ProvideMissionService() *MissionService {
	client := new(http.Client)

	return &MissionService{
		MissionsURL: "https://api.spacexdata.com/v3/missions",
		HTTPClient:  client,
	}
}

// AllMissions gathers all the registered missions
func (s *MissionService) AllMissions() ([]*entity.Mission, error) {
	resp, err := s.HTTPClient.Get(s.MissionsURL)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		bytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bytes)
		log.Println(bodyString)
	}

	return nil, nil
}
