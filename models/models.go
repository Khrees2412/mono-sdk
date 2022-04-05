package models

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
type UpdateAccountHolder struct {
	Status  string
	Message string
	Data    struct {
		Address struct {
			Address_line1 string
			Lga           string
			City          string
			State         string
		}
		Identity struct {
			Type   string
			Number string
			Url    string
		}
		Email string
		Bvn   string
	}
}

type VirtualAccount struct {
	Status  string
	Message string
}

type VirtualAccountRes struct {
	Status  string
	Message string
	Data    struct {
		Currency       string
		Balance        int
		Status         string
		App            string
		Business       string
		Id             string
		Bank_name      string
		Bank_code      string
		Kyc_level      string
		Account_name   string
		Account_number string
	}
}
type Account struct {
	Status  string
	Message string
	Data    struct {
		Address struct {
			Country       string
			State         string
			Lga           string
			City          string
			Address_line1 string
			Postal_code   string
		}
		Phone      string
		Entity     string
		Business   string
		App        string
		Live       bool
		Created_at string
		Updated_at string
		Id         string
		Identity   struct {
			Type   string
			Number string
			Url    string
		}
		First_name       string
		Last_name        string
		Selfie_url       string
		Virtual_accounts []string
	}
}
type Accounts struct {
	Status  string
	Message string
	Data    []struct {
		Address struct {
			Country       string
			State         string
			Lga           string
			City          string
			Address_line1 string
			Postal_code   string
		}
		Phone      string
		Entity     string
		Business   string
		App        string
		Live       bool
		Created_at string
		Updated_at string
		Id         string
		Identity   struct {
			Type   string
			Number string
			Url    string
		}
		First_name string
		Last_name  string
		Selfie_url string
	}
}
type AccountTxn struct {
	Status  string
	Message string
	Data    []struct {
		Currency string
		Amount   int
		Type     string
		Entry    string
		Date     string
		Business struct {
			Name string
			Id   string
		}
		Id string
	}
	Meta struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}
type Txn struct {
	Status  string
	Message string
	Data    struct {
		Currency  string
		Amount    int
		Type      string
		Entry     string
		Date      string
		Narration string
		Business  struct {
			Name string
			Id   string
		}
		Id string
	}
	Meta struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}
type Txns struct {
	Status  string
	Message string
	Data    []struct {
		Id             string
		Currency       string
		Amount         int
		Type           string
		Entry          string
		Date           string
		Account        string
		Account_holder string
		Narration      string
		Business       struct {
			Name string
			Id   string
		}
	}
	Meta struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}

type Budget struct {
	Status  string
	Message string
}

type VirtualCard struct {
}
