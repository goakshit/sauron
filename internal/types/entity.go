package types

type MerchantDetails struct {
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Perc  float64 `json:"perc"`
}

type UserDetails struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	DueAmount   float64 `json:"due_amount"`
	CreditLimit float64 `json:"credit_limit"`
}
