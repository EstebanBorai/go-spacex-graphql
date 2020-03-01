package entity

// Mission represents an SpaceX Mission
type Mission struct {
	MissionName   string   `json:"mission_name"`
	MissionID     string   `json:"mission_id"`
	Manufacturers []string `json:"manufacturers"`
	PayloadIDs    []string `json:"payload_ids"`
	Wikipedia     string   `json:"wikipedia"`
	Website       string   `json:"website"`
	Twitter       string   `json:"twitter"`
	Description   string   `json:"dscription"`
}
