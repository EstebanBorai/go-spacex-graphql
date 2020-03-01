package resolver

// Mission resolves Mission's Queries
func (r *Resolver) Mission() ([]*MissionResolver, error) {
	var missionsResolvers []*MissionResolver = make([]*MissionResolver, 0)

	response, err := r.services.MissionService.AllMissions()

	if err != nil {
		return missionsResolvers, err
	}

	for _, mission := range response {
		missionsResolvers = append(missionsResolvers, &MissionResolver{
			m: mission,
		})
	}

	return missionsResolvers, nil
}
