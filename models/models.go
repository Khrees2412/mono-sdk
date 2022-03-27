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
			Type     string
			// PERSONAL_BANKING or BUSINESS_BANKING
		}
		Name          string
		AccountNumber string
		Type          string
		Balance       int
		Currency      string
		BVN           string
	}
}

type Asset struct {
	Id           string
	Name         string
	AssestType   string
	Cost         int
	Return       int
	Quantity     int
	Currency     string
	AssetDetails struct {
		symbol string
		price  int
	}
}

type Earnings struct {
	Id        string
	Amount    int
	Narration string
	Date      string
	Asset     struct {
		symbol        string
		Name          string
		Sale_price    int
		Quantity_sold int
	}
}
type Identity struct {
	FullName      string
	Email         string
	Phone         string
	Gender        string
	DOB           string
	BVN           string
	MaritalStatus string
	AddressLine1  string
	AddressLine2  string
}
type Income struct {
	Average_income int
	Monthly_income int
	Yearly_income  int
	Income_sources int
}
