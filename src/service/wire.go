// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package service

import (
	"github.com/google/wire"
)

// ServiceProvider contains avilable services
type ServiceProvider struct {
	CapsuleService *CapsuleService
	MissionService *MissionService
}

// ProvideServiceProvider creates an instance of ServiceProvider
func ProvideServiceProvider(capsuleService *CapsuleService, missionService *MissionService) ServiceProvider {
	return ServiceProvider{
		CapsuleService: capsuleService,
		MissionService: missionService,
	}
}

// InitProvider inject dependencies for services
// and provides a service provider
func InitProvider() ServiceProvider {
	wire.Build(
		ProvideCapsuleService,
		ProvideMissionService,
		ProvideServiceProvider,
	)

	return ServiceProvider{}
}
