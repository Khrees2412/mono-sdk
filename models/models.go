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

type VirtualAccountsRes struct {
	Status  string
	Message string
	Data    []struct {
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

type Wallet struct {
	Status  string
	Message string
	Data    struct {
		Link string
	}
}

type WalletRes struct {
	Status  string
	Message string
	Data    []struct {
		Currency  string
		Balance   int
		Threshold struct {
			Value   int
			Enabled bool
		}
	}
}

type VirtualAccountBalance struct {
	Status  string
	Message string
	Data    struct {
		Withdrawable_amount int
		Available_amount    int
	}
}

type VirtualAccountKyc struct {
	Status  string
	Message string
	Data    struct {
		Maximum_credit_amount         string
		Maximum_transfer_amount       string
		Maximum_daily_transfer_amount string
		Maximum_account_balance       string
	}
}
type VirtualAccountUpgrade struct {
	Status  string
	Message string
	Data    struct {
		Id        string
		Kyc_level string
	}
}
type MockTxn struct {
	Status  string
	Message string
	Data    struct {
		Id string
	}
}
type FundTransfer struct {
	Status  string
	Message string
	Data    struct {
		Id string
	}
}

type Cac struct {
	State            string
	Id               int
	Address          string
	ApprovedName     string
	RcNumber         string
	BranchAddress    string
	RegistrationDate string
	ClassificationId int
	Email            string
	Lga              string
	City             string
	Status           string
}

type Shareholders []struct {
	EntityType string
	Name       string
	Role       string
	Details    struct {
		Id             int
		Address        string
		City           string
		Residence      string
		Gender         string
		Nationality    string
		Occupation     string
		Shares         int
		Status         string
		Email          string
		Phone          string
		Dob            string
		IdentityType   string
		IdentityNumber string
	}
}
type VerifyAcct struct {
	Status  string
	Message string
	Data    struct {
		Account_number string
		Account_name   string
	}
}
