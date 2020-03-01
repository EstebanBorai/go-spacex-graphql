package resolver

import (
	"github.com/estebanborai/go-spacex-graphql/src/entity"
)

type MissionResolver struct {
	m *entity.Mission
}

func (r *MissionResolver) MissionName() string {
	return r.m.MissionName
}

func (r *MissionResolver) MissionID() string {
	return r.m.MissionID
}

func (r *MissionResolver) Manufacturers() []*string {
	var manufacturers []*string

	manufacturers = make([]*string, 0)

	for _, manufacturer := range r.m.Manufacturers {
		manufacturers = append(manufacturers, &manufacturer)
	}

	return manufacturers
}

func (r *MissionResolver) PayloadIDs() []*string {
	var payloadIDs []*string

	payloadIDs = make([]*string, 0)

	for _, payloadID := range r.m.PayloadIDs {
		payloadIDs = append(payloadIDs, &payloadID)
	}

	return payloadIDs
}

func (r *MissionResolver) Wikipedia() *string {
	return &r.m.Wikipedia
}

func (r *MissionResolver) Website() *string {
	return &r.m.Website
}

func (r *MissionResolver) Twitter() *string {
	return &r.m.Twitter
}

func (r *MissionResolver) Description() *string {
	return &r.m.Description
}
