package types

type BudgetCategoryTypeStats struct {
	CategoryType      string  `json:"category_type" db:"category_type"`
	BgColor           string  `json:"bg_color" db:"bgColor"`
	TextColor         string  `json:"text_color" db:"textColor"`
	Month             int     `json:"month" db:"month"`
	Year              int     `json:"year" db:"year"`
	ItemCount         int     `json:"item_count" db:"item_count"`
	TotalActualAmount float64 `json:"total_actual_amount" db:"total_actual_amount"`
	Categories        string  `json:"categories" db:"categories"`
}

type BudgetStatsByMonthAndYear struct {
	ItemID       string  `db:"itemId"`
	ItemName     string  `db:"itemName"`
	ActualAmount float64 `db:"actualAmount"`
	Status       string  `db:"status"`
	Month        int     `db:"month"`
	Year         int     `db:"year"`
	CategoryName string  `db:"categoryName"`
	Type         string  `db:"type"`
	BgColor      string  `db:"bgColor"`
	TextColor    string  `db:"textColor"`
}

type RawUserStats struct {
	Type         string  `json:"type"`
	CategoryName string  `json:"categoryName"`
	ItemName     string  `json:"itemName"`
	ActualAmount float64 `json:"actualAmount"`
	Month        int     `json:"month"`
	Year         int     `json:"year"`
}

type CopyBudgetReqBody struct {
	OldMonth     int `json:"oldMonth"`
	OldYear      int `json:"oldYear"`
	CurrentMonth int `json:"currentMonth"`
	CurrentYear  int `json:"currentYear"`
}

type BudgetCategoryReqBody struct {
	CategoryName        string `json:"categoryName"`
	CategoryDescription string `json:"categoryDescription"`
	Month               int    `json:"month"`
	Year                int    `json:"year"`
	ParentId            string `json:"parentId"`
	CategoryTypeId      string `json:"categoryTypeId"`
}

type BudgetCategoryTypeReqBody struct {
	Type      string `json:"type"`
	BgColor   string `json:"bgColor"`
	TextColor string `json:"textColor"`
}

type BudgetItemReqBody struct {
	CategoryID      string  `json:"categoryId"`
	ItemName        string  `json:"itemName"`
	Description     string  `json:"description,omitempty"`
	AllocatedAmount float64 `json:"allocatedAmount,omitempty"`
	ActualAmount    float64 `json:"actualAmount,omitempty"`
	CurrencyCode    string  `json:"currencyCode,omitempty"`
	Status          string  `json:"status,omitempty"`
	Month           int     `json:"month"`
	Year            int     `json:"year"`
}

type BudgetCategoryType struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	BgColor   string `json:"bgColor"`
	TextColor string `json:"textColor"`
}

type BudgetCategory struct {
	Id                  string `json:"id"`
	CategoryName        string `json:"categoryName"`
	CategoryDescription string `json:"categoryDescription"`
	Month               int    `json:"month"`
	Year                int    `json:"year"`
	ParentId            string `json:"parentId"`
	CategoryTypeId      string `json:"categoryTypeId"`
}

type BudgetItem struct {
	Id              string  `json:"id"`
	CategoryID      string  `json:"categoryId"`
	ItemName        string  `json:"itemName"`
	Description     string  `json:"description"`
	AllocatedAmount float64 `json:"allocatedAmount"`
	ActualAmount    float64 `json:"actualAmount"`
	CurrencyCode    string  `json:"currencyCode"`
	Status          string  `json:"status"`
	Month           int     `json:"month"`
	Year            int     `json:"year"`
}
