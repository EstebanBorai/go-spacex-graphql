package resolver

// Capsule resolves Capsule's Queries
func (r *Resolver) Capsule() ([]*CapsuleResolver, error) {
	var capsulesResolvers []*CapsuleResolver = make([]*CapsuleResolver, 0)

	response, err := r.services.CapsuleService.AllCapsules()

	if err != nil {
		return capsulesResolvers, err
	}

	for _, capsule := range response {
		capsulesResolvers = append(capsulesResolvers, &CapsuleResolver{
			c: capsule,
		})
	}

	return capsulesResolvers, nil
}
