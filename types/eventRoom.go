package types

type EventRoomRequestBody struct {
	Name               string `json:"name"`               // Maps to "name" in JSON
	Capacity           string `json:"capacity"`           // Maps to "capacity" in JSON
	Location           string `json:"location"`           // Maps to "location" in JSON
	IsUnderMaintenance bool   `json:"isUnderMaintenance"` // Maps to "isUnderMaintenance" in JSON
	StartTime          string `json:"startTime"`          // Maps to "startTime" in JSON
	EndTime            string `json:"endTime"`            // Maps to "endTime" in JSON
}

type EventRoom struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`               // Maps to "name" in JSON
	Capacity           string  `json:"capacity"`           // Maps to "capacity" in JSON
	Location           string  `json:"location"`           // Maps to "location" in JSON
	IsUnderMaintenance bool    `json:"isUnderMaintenance"` // Maps to "isUnderMaintenance" in JSON
	StartTime          string  `json:"startTime"`          // Maps to "startTime" in JSON
	EndTime            string  `json:"endTime"`            // Maps to "endTime" in JSON
	CompanyId          string  `json:"companyId"`
	CreatedAt          []uint8 `json:"created_at" db:"created_at"`
	UpdatedAt          []uint8 `json:"updated_at" db:"updated_at"`
}
