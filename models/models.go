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

type Transaction struct {
	Paging struct {
		Total    int
		Page     int
		Previous string
		Next     string
	}
	Data []struct {
		Id        string
		Amount    int
		Date      string
		Narration string
		Type      string
		Category  string
	}
}

type StatementJSON struct {
	Meta struct {
		count int
	}
	Data []struct {
		Id        string
		Amount    string
		Date      string
		Narration string
		Type      string
		Category  string
	}
}
type StatementPDF struct {
	Id     string
	Status string
	Path   string
}

type View360 struct {
	Message string
	Data    []struct {
		AccountNumber string
		Institution   struct {
			Id       string
			Name     string
			BankCode string
			Icon     string
		}
	}
}
type Sync struct {
	Status     string
	Code       string
	HasNewData bool
}

type Reauth struct {
	Token string
}

type MakePayment struct {
	Id           string
	Type         string
	Amount       int
	Description  string
	Reference    string
	Payment_link string
	Created_at   string
	Updated_at   string
}

type VerifyPayment struct {
	Type string
	Data struct {
		_Id         string
		Id          string
		Status      string
		Amount      int
		Description string
		Fee         int
		Currency    string
		Account     string
		Customer    string
		Reference   string
		Created_at  string
		Updated_at  string
	}
}

type GetPayment struct {
	Payments []struct {
		_Id         string
		Id          string
		Status      string
		Amount      int
		Description string
		Currency    string
		Account     struct {
			_Id         string
			Institution struct {
				Id   string
				Name string
				Type string
				// PERSONAL_BANKING or BUSINESS_BANKING
			}
			Name          string
			Currency      string
			AccountNumber string
			Created_at    string
			Updated_at    string
		}
		Customer   string
		Reference  string
		Created_at string
		Updated_at string
	}
	Paging struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}

type AccountHolder struct {
	Status  string
	Message string
	Data    struct {
		Id string
	}
}

type Budget struct {
	Status  string
	Message string
}
