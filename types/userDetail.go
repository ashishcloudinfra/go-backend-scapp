package types

type MemberShortDetail struct {
	ID        string  `json:"id" db:"id"`
	FirstName string  `json:"firstName" db:"firstName"`
	LastName  string  `json:"lastName" db:"lastName"`
	Gender    *string `json:"gender,omitempty" db:"gender"`
	Status    string  `json:"status"`
	Phone     string  `json:"phone" db:"phone"`
	Email     string  `json:"email" db:"email"`
	IamID     string  `json:"iamId" db:"iamId"`
}

type PersonalDetail struct {
	ID        string  `json:"id" db:"id"`
	FirstName string  `json:"firstName" db:"firstName"`
	LastName  string  `json:"lastName" db:"lastName"`
	Email     string  `json:"email" db:"email"`
	Phone     string  `json:"phone" db:"phone"`
	Address   *string `json:"address,omitempty" db:"address"`
	City      *string `json:"city,omitempty" db:"city"`
	State     *string `json:"state,omitempty" db:"state"`
	Zip       *string `json:"zip,omitempty" db:"zip"`
	Country   *string `json:"country,omitempty" db:"country"`
	Dob       *string `json:"dob,omitempty" db:"dob"`
	Gender    *string `json:"gender,omitempty" db:"gender"`
	Metadata  *string `json:"metadata,omitempty" db:"metadata"`
	CreatedAt []uint8 `json:"created_at" db:"created_at"`
	UpdatedAt []uint8 `json:"updated_at" db:"updated_at"`
	IamID     string  `json:"iamId" db:"iamId"`
}

type PersonalDetailWithStatus struct {
	ID          string  `json:"id" db:"id"`
	FirstName   string  `json:"firstName" db:"firstName"`
	LastName    string  `json:"lastName" db:"lastName"`
	Email       string  `json:"email" db:"email"`
	Phone       string  `json:"phone" db:"phone"`
	Address     string  `json:"address,omitempty" db:"address"`
	City        string  `json:"city,omitempty" db:"city"`
	State       string  `json:"state,omitempty" db:"state"`
	Zip         string  `json:"zip,omitempty" db:"zip"`
	Country     string  `json:"country,omitempty" db:"country"`
	Dob         string  `json:"dob,omitempty" db:"dob"`
	Gender      string  `json:"gender,omitempty" db:"gender"`
	Metadata    string  `json:"metadata,omitempty" db:"metadata"`
	CreatedAt   []uint8 `json:"created_at" db:"created_at"`
	UpdatedAt   []uint8 `json:"updated_at" db:"updated_at"`
	IamID       string  `json:"iamId" db:"iamId"`
	Status      string  `json:"status"`
	Role        string  `json:"role"`
	Permissions string  `json:"permissions"`
	Plans       string  `json:"plans"`
}

type UserDetailPostRequestBody struct {
	FirstName   string `json:"firstName" db:"firstName"`
	LastName    string `json:"lastName" db:"lastName"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	Address     string `json:"address,omitempty" db:"address"`
	City        string `json:"city,omitempty" db:"city"`
	State       string `json:"state,omitempty" db:"state"`
	Zip         string `json:"zip,omitempty" db:"zip"`
	Country     string `json:"country,omitempty" db:"country"`
	Dob         string `json:"dob,omitempty" db:"dob"`
	Gender      string `json:"gender,omitempty" db:"gender"`
	Metadata    string `json:"metadata,omitempty" db:"metadata"`
	Role        string `json:"role"`
	Permissions string `json:"permissions"`
	Plans       string `json:"plans"`
	Status      string `json:"status"`
}

type UserSignUpRequestBody struct {
	FirstName   string `json:"firstName" db:"firstName"`
	LastName    string `json:"lastName" db:"lastName"`
	Email       string `json:"email" db:"email"`
	Phone       string `json:"phone" db:"phone"`
	Address     string `json:"address,omitempty" db:"address"`
	City        string `json:"city,omitempty" db:"city"`
	State       string `json:"state,omitempty" db:"state"`
	Zip         string `json:"zip,omitempty" db:"zip"`
	Country     string `json:"country,omitempty" db:"country"`
	Dob         string `json:"dob,omitempty" db:"dob"`
	Gender      string `json:"gender,omitempty" db:"gender"`
	Metadata    string `json:"metadata,omitempty" db:"metadata"`
	Role        string `json:"role"`
	Permissions string `json:"permissions"`
	Plans       string `json:"plans"`
	Status      string `json:"status"`
	UserName    string `json:"userName"`
	Password    string `json:"password"`
}

type UserDetailReq struct {
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName" db:"lastName"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
	Address   string `json:"address,omitempty" db:"address"`
	City      string `json:"city,omitempty" db:"city"`
	State     string `json:"state,omitempty" db:"state"`
	Zip       string `json:"zip,omitempty" db:"zip"`
	Country   string `json:"country,omitempty" db:"country"`
	Dob       string `json:"dob,omitempty" db:"dob"`
	Gender    string `json:"gender,omitempty" db:"gender"`
	Metadata  string `json:"metadata,omitempty" db:"metadata"`
}
