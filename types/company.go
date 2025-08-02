package types

type PostCompanyRequestBody struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Contact   string `json:"contact"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Pincode   string `json:"pincode"`
	Type      string `json:"type"`
	IsDefault bool   `json:"isDefault"`
	OwnerID   string `json:"ownerId"`
}

type Company struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Contact   string  `json:"contact"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Pincode   string  `json:"pincode"`
	Type      string  `json:"type"`
	IsDefault bool    `json:"isDefault"`
	IamId     string  `json:"iamId"`
	CreatedAt []uint8 `json:"created_at" db:"created_at"`
	UpdatedAt []uint8 `json:"updated_at" db:"updated_at"`
}
