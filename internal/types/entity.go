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

type TxnDetails struct {
	ID           int     `json:"id" gorm:"primaryKey"`
	UserName     string  `json:"user_name"`
	MerchantName string  `json:"merchant_name"`
	MerchantPerc float64 `json:"merchant_perc"`
	Amount       float64 `json:"amount"`
}

type PaybackDetails struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
}

type ReportUserDues struct {
	Name      string  `json:"name"`
	DueAmount float64 `json:"due_amount"`
}

type ReportMerchantTxn struct {
	Amount       float64 `json:"amount"`
	MerchantPerc float64 `json:"merchant_perc"`
}
