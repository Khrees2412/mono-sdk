package models

type Details struct {
	Meta struct {
		// Available, Processing, Failed
		DataStatus string `json:"data_status"`
		// mobile_banking, internet_banking
		AuthMethod string `json:"auth_method"`
	}
	Account struct {
		Id          string `json:"id" :"id"`
		Institution struct {
			Name     string `json:"name,omitempty"`
			BankCode string `json:"bank_code,omitempty"`
			// PERSONAL_BANKING or BUSINESS_BANKING
			Type string `json:"type" json:"type,omitempty"`
		} `json:"institution"`
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
		Type          string `json:"type"`
		Balance       int    `json:"balance"`
		Currency      string `json:"currency"`
		BVN           string `json:"bvn"`
	}
}

type Asset struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	AssetType    string `json:"asset_type,omitempty"`
	Cost         int    `json:"cost,omitempty"`
	Return       int    `json:"return,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
	Currency     string `json:"currency,omitempty"`
	AssetDetails struct {
		symbol string `json:"symbol,omitempty"`
		price  int    `json:"price,omitempty"`
	} `json:"asset_details"`
}

type Earnings struct {
	Id        string `json:"id,omitempty"`
	Amount    int    `json:"amount,omitempty"`
	Narration string `json:"narration,omitempty"`
	Date      string `json:"date,omitempty"`
	Asset     struct {
		symbol       string `json:"symbol,omitempty"`
		Name         string `json:"name,omitempty"`
		SalePrice    int    `json:"sale_price,omitempty"`
		QuantitySold int    `json:"quantity_sold,omitempty"`
	} `json:"asset"`
}
type Identity struct {
	FullName      string `json:"full_name,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Gender        string `json:"gender,omitempty"`
	DOB           string `json:"dob,omitempty"`
	BVN           string `json:"bvn,omitempty"`
	MaritalStatus string `json:"marital_status,omitempty"`
	AddressLine1  string `json:"address_line1,omitempty"`
	AddressLine2  string `json:"address_line2,omitempty"`
}
type Income struct {
	AverageIncome int `json:"average_income,omitempty"`
	MonthlyIncome int `json:"monthly_income,omitempty"`
	YearlyIncome  int `json:"yearly_income,omitempty"`
	IncomeSources int `json:"income_sources,omitempty"`
}

type Transaction struct {
	Paging struct {
		Total    int    `json:"total,omitempty"`
		Page     int    `json:"page,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"paging"`
	Data []struct {
		Id        string `json:"id,omitempty"`
		Amount    int    `json:"amount,omitempty"`
		Date      string `json:"date,omitempty"`
		Narration string `json:"narration,omitempty"`
		Type      string `json:"type,omitempty"`
		Category  string `json:"category,omitempty"`
	} `json:"data"`
}

type StatementJSON struct {
	Meta struct {
		count int `json:"count,omitempty"`
	} `json:"meta"`
	Data []struct {
		Id        string `json:"id,omitempty"`
		Amount    string `json:"amount,omitempty"`
		Date      string `json:"date,omitempty"`
		Narration string `json:"narration,omitempty"`
		Type      string `json:"type,omitempty"`
		Category  string `json:"category,omitempty"`
	} `json:"data"`
}
type StatementPDF struct {
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Path   string `json:"path,omitempty"`
}

type View360 struct {
	Message string `json:"message,omitempty"`
	Data    []struct {
		AccountNumber string `json:"account_number,omitempty"`
		Institution   struct {
			Id       string `json:"id,omitempty"`
			Name     string `json:"name,omitempty"`
			BankCode string `json:"bank_code,omitempty"`
			Icon     string `json:"icon,omitempty"`
		} `json:"institution"`
	} `json:"data,omitempty"`
}
type Sync struct {
	Status     string `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	HasNewData bool   `json:"has_new_data,omitempty"`
}

type Reauth struct {
	Token string `json:"token,omitempty"`
}
