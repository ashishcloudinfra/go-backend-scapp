package types

type Variety struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type MenuItemRequestBody struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Photo       *string   `json:"photo,omitempty"` // Optional field
	IsVeg       bool      `json:"isVeg"`           // Veg or Non-Veg
	Category    string    `json:"category"`        // Indian, Italian, Starters
	CookingTime string    `json:"cookingTime"`
	Varieties   []Variety `json:"varieties"` // Array of varieties
}

type MenuItem struct {
	MenuItemID          string `json:"menuItemId"`
	MenuItemName        string `json:"menuItemName"`
	MenuItemDescription string `json:"menuItemDescription"`
	CookingTime         string `json:"cookingTime"`
	MenuItemPhoto       string `json:"menuItemPhoto"`
	IsVeg               bool   `json:"isVeg"`
	CategoryName        string `json:"categoryName"`
	PricingID           string `json:"pricingId"`
	VarietyType         string `json:"varietyType"`
	Price               string `json:"price"`
}
