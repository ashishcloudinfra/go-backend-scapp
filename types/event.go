package types

type MetadataRecurrenceType struct {
	Every           *string  `json:"every,omitempty"`           // Optional field
	RepeatEvery     *int     `json:"repeatEvery,omitempty"`     // Optional field
	RepeatOn        []string `json:"repeatOn,omitempty"`        // Optional field
	MonthDate       *int     `json:"monthDate,omitempty"`       // Optional field
	MonthWeekNumber *string  `json:"monthWeekNumber,omitempty"` // Optional field
	SelectedWeekDay *string  `json:"selectedWeekDay,omitempty"` // Optional field
	SelectorType    *string  `json:"selectorType,omitempty"`    // Optional field
}

type RecurrenceType string

const (
	Daily   RecurrenceType = "daily"
	Weekly  RecurrenceType = "weekly"
	Monthly RecurrenceType = "monthly"
	Yearly  RecurrenceType = "yearly"
)

type EventMetadata struct {
	Daily   MetadataRecurrenceType `json:"daily"`
	Weekly  MetadataRecurrenceType `json:"weekly"`
	Monthly MetadataRecurrenceType `json:"monthly"`
	Yearly  MetadataRecurrenceType `json:"yearly"`
}

type Status string

const (
	Authorized Status = "authorized"
	Review     Status = "review"
)

type EventFormValues struct {
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	OrganiserId    string        `json:"organiserId"`
	EventRoomId    string        `json:"eventRoomId"`
	StartDate      string        `json:"startDate"`
	EndDate        string        `json:"endDate"`
	StartTime      string        `json:"startTime"`
	EndTime        string        `json:"endTime"`
	IsAllDayEvent  bool          `json:"isAllDayEvent"`
	IsRecurring    bool          `json:"isRecurring"`
	RecurrenceType string        `json:"recurrenceType"`
	Status         Status        `json:"status"`
	Metadata       EventMetadata `json:"metadata"`
}

type Event struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Description    *string       `json:"description,omitempty"`
	OrganiserID    string        `json:"organiserId"`
	EventRoomID    *string       `json:"eventRoomId,omitempty"`
	StartDate      string        `json:"startDate"`
	EndDate        string        `json:"endDate"`
	StartTime      string        `json:"startTime"`
	EndTime        string        `json:"endTime"`
	IsAllDayEvent  bool          `json:"isAllDayEvent"`
	IsRecurring    bool          `json:"isRecurring"`
	RecurrenceType *string       `json:"recurrenceType,omitempty"`
	Status         string        `json:"status"`
	Metadata       EventMetadata `json:"metadata"`
	CreatedAt      []uint8       `json:"created_at" db:"created_at"`
	UpdatedAt      []uint8       `json:"updated_at" db:"updated_at"`
	CompanyID      string        `json:"companyId"`
}

type EventWithOrganiserAndRoomName struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Description    *string       `json:"description,omitempty"`
	OrganiserID    string        `json:"organiserId"`
	EventRoomID    *string       `json:"eventRoomId,omitempty"`
	StartDate      string        `json:"startDate"`
	EndDate        string        `json:"endDate"`
	StartTime      string        `json:"startTime"`
	EndTime        string        `json:"endTime"`
	IsAllDayEvent  bool          `json:"isAllDayEvent"`
	IsRecurring    bool          `json:"isRecurring"`
	RecurrenceType *string       `json:"recurrenceType,omitempty"`
	Status         string        `json:"status"`
	Metadata       EventMetadata `json:"metadata"`
	CreatedAt      []uint8       `json:"created_at" db:"created_at"`
	UpdatedAt      []uint8       `json:"updated_at" db:"updated_at"`
	CompanyID      string        `json:"companyId"`
	Organiser      string        `json:"organiser"`
	EventRoomName  *string       `json:"eventRoomName,omitempty"`
}
