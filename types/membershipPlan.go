package types

type MembershipPlanFormValues struct {
	Name               string   `json:"name"`
	Duration           string   `json:"duration"`
	Features           []string `json:"features"`
	CancellationPolicy []string `json:"cancellation_policy"`
	Price              float64  `json:"price"`
	Discount           string   `json:"discount"`
	Status             string   `json:"status"`
}

type MembershipPlan struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Price              float64  `json:"price"` // NUMERIC(10, 2) maps to float64
	Duration           string   `json:"duration"`
	Discount           *string  `json:"discount,omitempty"` // Nullable field
	Features           []string `json:"features"`           // TEXT[] maps to []string
	Status             string   `json:"status"`
	CancellationPolicy []string `json:"cancellation_policy"` // TEXT[] maps to []string
	CompanyId          string   `json:"companyId"`           // Foreign key
	CreatedAt          []uint8  `json:"created_at" db:"created_at"`
	UpdatedAt          []uint8  `json:"updated_at" db:"updated_at"`
}
