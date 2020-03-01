package entity

import "time"

// Capsule represents an SpaceX Rocket Capsule
// properties
type Capsule struct {
	CapsuleSerial      string    `json:"capsule_serial"`
	CapsuleID          string    `json:"capsule_id"`
	Status             string    `json:"status"`
	OriginalLaunch     time.Time `json:"original_launch"`
	OriginalLaunchUnix int64     `json:"original_launch_unix"`
	Missions           []Mission `json:"missions"`
	Landings           int32     `json:"landings"`
	Type               string    `json:"type"`
	Details            string    `json:"details"`
	ReuseCount         int16     `json:"reuse_count"`
}
