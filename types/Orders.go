package types

type OrderItemRequestBody struct {
	Quantity          int    `json:"quantity"`
	MenuItemPricingId string `json:"menuItemPricingId"`
}

type OrderRequestBody struct {
	TableNumber string                 `json:"tableNumber"`
	IsDineIn    bool                   `json:"isDineIn"`
	Status      string                 `json:"status"`
	Phone       string                 `json:"phone"`
	Email       string                 `json:"email"`
	Items       []OrderItemRequestBody `json:"items"`
	CompanyID   string                 `json:"companyId"`
}

type MenuWithQuantity struct {
	Quantity     int
	MenuitemName string
	VarietyType  string
	Price        string
	Photo        string
	IsVeg        bool
}

type AdminOrder struct {
	OrderId      string
	OrderItemId  string
	OrderNumber  int
	Status       string
	MenuitemName string
	VarietyType  string
	IsVeg        bool
	Quantity     int
	TableNumber  string
}
