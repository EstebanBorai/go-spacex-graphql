package resolver

import (
	"github.com/estebanborai/go-spacex-graphql/src/entity"
	graphql "github.com/graph-gophers/graphql-go"
)

type CapsuleResolver struct {
	c *entity.Capsule
}

func (r *CapsuleResolver) CapsuleSerial() string {
	return r.c.CapsuleSerial
}

func (r *CapsuleResolver) CapsuleID() string {
	return r.c.CapsuleID
}

func (r *CapsuleResolver) Status() string {
	return r.c.Status
}

func (r *CapsuleResolver) OriginalLaunch() graphql.Time {
	return graphql.Time{
		Time: r.c.OriginalLaunch,
	}
}

func (r *CapsuleResolver) OriginalLaunchUnix() int32 {
	return int32(r.c.OriginalLaunchUnix)
}

func (r *CapsuleResolver) Missions() []*MissionResolver {
	var missionResolvers []*MissionResolver

	missionResolvers = make([]*MissionResolver, 0)

	for _, mission := range r.c.Missions {
		missionResolvers = append(missionResolvers, &MissionResolver{
			m: &mission,
		})
	}

	return missionResolvers
}

func (r *CapsuleResolver) Landings() int32 {
	return r.c.Landings
}

func (r *CapsuleResolver) Type() string {
	return r.c.Type
}

func (r *CapsuleResolver) Details() *string {
	return &r.c.Details
}

func (r *CapsuleResolver) ReuseCount() *int32 {
	num := int32(r.c.ReuseCount)

	return &num
}
