package types

import (
	"regexp"
	"time"
)

type IAM struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type IAMRequestBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Contact   string `json:"contact"`
}

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupRequestBody struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	Contact     string `json:"contact"`
	Permissions string `json:"permissions"`
	Plans       string `json:"plans"`
	CompanyId   string `json:"companyId,omitempty"`
}

// BypassRule defines a rule for bypassing JWT validation.
// Pattern is the regex for the route, and Methods is the list of HTTP methods
// for which this rule applies. If Methods is empty, the rule applies to all methods.
type BypassRule struct {
	Pattern *regexp.Regexp
	Methods []string
}
