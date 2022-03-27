package models

type Details struct {
	Meta struct {
		Data_status string
		// Available, Processing, Failed
		Auth_method string
		// mobile_banking, internet_banking
	}
	Account struct {
		Id          string
		Institution struct {
			Name     string
			BankCode string
			BankType string
			// PERSONAL_BANKING or BUSINESS_BANKING
		}
		Name          string
		AccountNumber string
		AccountType   string
		Balance       int
		Currency      string
		BVN           string
	}
}
