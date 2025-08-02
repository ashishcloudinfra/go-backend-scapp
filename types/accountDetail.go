package types

type AccountDetail struct {
	ID          string  `json:"id" db:"id"`
	Role        string  `json:"role" db:"role"`
	Permissions string  `json:"permissions" db:"permissions"`
	Plans       string  `json:"plans" db:"plans"`
	Status      string  `json:"status" db:"status"`
	CreatedAt   []uint8 `json:"created_at" db:"created_at"`
	UpdatedAt   []uint8 `json:"updated_at" db:"updated_at"`
	CompanyID   string  `json:"companyId" db:"companyId"`
	IamID       string  `json:"iamId" db:"iamId"`
}
