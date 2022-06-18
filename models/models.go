package models

type MakePayment struct {
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Reference   string `json:"reference,omitempty"`
	PaymentLink string `json:"payment_link,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type VerifyPayment struct {
	Type string `json:"type,omitempty"`
	Data struct {
		_Id         string `json:"_id,omitempty"`
		Id          string `json:"id,omitempty"`
		Status      string `json:"status,omitempty"`
		Amount      int    `json:"amount,omitempty"`
		Description string `json:"description,omitempty"`
		Fee         int    `json:"fee,omitempty"`
		Currency    string `json:"currency,omitempty"`
		Account     string `json:"account,omitempty"`
		Customer    string `json:"customer,omitempty"`
		Reference   string `json:"reference,omitempty"`
		CreatedAt   string `json:"created___at,omitempty"`
		UpdatedAt   string `json:"updated___at,omitempty"`
	} `json:"data"`
}

type GetPayment struct {
	Payments []struct {
		_Id         string `json:"_id,omitempty"`
		Id          string `json:"id,omitempty"`
		Status      string `json:"status,omitempty"`
		Amount      int    `json:"amount,omitempty"`
		Description string `json:"description,omitempty"`
		Currency    string `json:"currency,omitempty"`
		Account     struct {
			_Id         string `json:"_id,omitempty"`
			Institution struct {
				Id   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
				// PERSONAL_BANKING or BUSINESS_BANKING
				Type string `json:"type,omitempty"`
			} `json:"institution"`
			Name          string `json:"name,omitempty"`
			Currency      string `json:"currency,omitempty"`
			AccountNumber string `json:"account_number,omitempty"`
			CreatedAt     string `json:"created_at,omitempty"`
			UpdatedAt     string `json:"updated_at,omitempty"`
		} `json:"account"`
		Customer  string `json:"customer,omitempty"`
		Reference string `json:"reference,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	} `json:"payments"`
	Paging struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	}
}

type AccountHolder struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Id string `json:"id,omitempty"`
	} `json:"data"`
}
type UpdateAccountHolder struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Address struct {
			AddressLine1 string `json:"address_line1,omitempty"`
			Lga          string `json:"lga,omitempty"`
			City         string `json:"city,omitempty"`
			State        string `json:"state,omitempty"`
		} `json:"address"`
		Identity struct {
			Type   string `json:"type,omitempty"`
			Number string `json:"number,omitempty"`
			Url    string `json:"url,omitempty"`
		} `json:"identity"`
		Email string `json:"email,omitempty"`
		Bvn   string `json:"bvn,omitempty"`
	} `json:"data"`
}

type VirtualAccount struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type VirtualAccountRes struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Currency      string `json:"currency,omitempty"`
		Balance       int    `json:"balance,omitempty"`
		Status        string `json:"status,omitempty"`
		App           string `json:"app,omitempty"`
		Business      string `json:"business,omitempty"`
		Id            string `json:"id,omitempty"`
		BankName      string `json:"bank_name,omitempty"`
		BankCode      string `json:"bank_code,omitempty"`
		KycLevel      string `json:"kyc_level,omitempty"`
		AccountName   string `json:"account_name,omitempty"`
		AccountNumber string `json:"account_number,omitempty"`
	} `json:"data"`
}

type VirtualAccountsRes struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Currency      string `json:"currency,omitempty"`
		Balance       int    `json:"balance,omitempty"`
		Status        string `json:"status,omitempty"`
		App           string `json:"app,omitempty"`
		Business      string `json:"business,omitempty"`
		Id            string `json:"id,omitempty"`
		BankName      string `json:"bank_name,omitempty"`
		BankCode      string `json:"bank_code,omitempty"`
		KycLevel      string `json:"kyc_level,omitempty"`
		AccountName   string `json:"account_name,omitempty"`
		AccountNumber string `json:"account_number,omitempty"`
	} `json:"data,omitempty"`
}
type Account struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Address struct {
			Country      string `json:"country,omitempty"`
			State        string `json:"state,omitempty"`
			Lga          string `json:"lga,omitempty"`
			City         string `json:"city,omitempty"`
			AddressLine1 string `json:"address_line1,omitempty"`
			PostalCode   string `json:"postal_code,omitempty"`
		} `json:"address"`
		Phone     string `json:"phone,omitempty"`
		Entity    string `json:"entity,omitempty"`
		Business  string `json:"business,omitempty"`
		App       string `json:"app,omitempty"`
		Live      bool   `json:"live,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		Id        string `json:"id,omitempty"`
		Identity  struct {
			Type   string `json:"type,omitempty"`
			Number string `json:"number,omitempty"`
			Url    string `json:"url,omitempty"`
		} `json:"identity"`
		FirstName       string   `json:"first_name,omitempty"`
		LastName        string   `json:"last_name,omitempty"`
		SelfieUrl       string   `json:"selfie_url,omitempty"`
		VirtualAccounts []string `json:"virtual_accounts,omitempty"`
	} `json:"data"`
}
type Accounts struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Address struct {
			Country      string `json:"country,omitempty"`
			State        string `json:"state,omitempty"`
			Lga          string `json:"lga,omitempty"`
			City         string `json:"city,omitempty"`
			AddressLine1 string `json:"address_line_1,omitempty"`
			PostalCode   string `json:"postal_code,omitempty"`
		} `json:"address"`
		Phone     string `json:"phone,omitempty"`
		Entity    string `json:"entity,omitempty"`
		Business  string `json:"business,omitempty"`
		App       string `json:"app,omitempty"`
		Live      bool   `json:"live,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
		Id        string `json:"id,omitempty"`
		Identity  struct {
			Type   string `json:"type,omitempty"`
			Number string `json:"number,omitempty"`
			Url    string `json:"url,omitempty"`
		} `json:"identity"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		SelfieUrl string `json:"selfie_url,omitempty"`
	} `json:"data,omitempty"`
}
type AccountTxn struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Currency string `json:"currency,omitempty"`
		Amount   int    `json:"amount,omitempty"`
		Type     string `json:"type,omitempty"`
		Entry    string `json:"entry,omitempty"`
		Date     string `json:"date,omitempty"`
		Business struct {
			Name string `json:"name,omitempty"`
			Id   string `json:"id,omitempty"`
		} `json:"business"`
		Id string `json:"id,omitempty"`
	} `json:"data,omitempty"`
	Meta struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"meta"`
}
type Txn struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Currency  string `json:"currency,omitempty"`
		Amount    int    `json:"amount,omitempty"`
		Type      string `json:"type,omitempty"`
		Entry     string `json:"entry,omitempty"`
		Date      string `json:"date,omitempty"`
		Narration string `json:"narration,omitempty"`
		Business  struct {
			Name string `json:"name,omitempty"`
			Id   string `json:"id,omitempty"`
		} `json:"business"`
		Id string `json:"id,omitempty"`
	} `json:"data"`
	Meta struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"meta"`
}
type Txns struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Id            string `json:"id,omitempty"`
		Currency      string `json:"currency,omitempty"`
		Amount        int    `json:"amount,omitempty"`
		Type          string `json:"type,omitempty"`
		Entry         string `json:"entry,omitempty"`
		Date          string `json:"date,omitempty"`
		Account       string `json:"account,omitempty"`
		AccountHolder string `json:"account_holder,omitempty"`
		Narration     string `json:"narration,omitempty"`
		Business      struct {
			Name string `json:"name,omitempty"`
			Id   string `json:"id,omitempty"`
		} `json:"business"`
	} `json:"data,omitempty"`
	Meta struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"meta"`
}

type Budget struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type Wallet struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Link string `json:"link,omitempty"`
	} `json:"data"`
}

type WalletRes struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Currency  string `json:"currency,omitempty"`
		Balance   int    `json:"balance,omitempty"`
		Threshold struct {
			Value   int  `json:"value,omitempty"`
			Enabled bool `json:"enabled,omitempty"`
		} `json:"threshold"`
	} `json:"data,omitempty"`
}

type VirtualAccountBalance struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		WithdrawableAmount int `json:"withdrawable_amount,omitempty"`
		AvailableAmount    int `json:"available_amount,omitempty"`
	} `json:"data"`
}

type VirtualAccountKyc struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		MaximumCreditAmount        string `json:"maximum_credit_amount,omitempty"`
		MaximumTransferAmount      string `json:"maximum_transfer_amount,omitempty"`
		MaximumDailyTransferAmount string `json:"maximum_daily_transfer_amount,omitempty"`
		MaximumAccountBalance      string `json:"maximum_account_balance,omitempty"`
	} `json:"data"`
}
type VirtualAccountUpgrade struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Id       string `json:"id,omitempty"`
		KycLevel string `json:"kyc_level,omitempty"`
	} `json:"data"`
}
type MockTxn struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Id string `json:"id,omitempty"`
	} `json:"data"`
}
type FundTransfer struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Id string `json:"id,omitempty"`
	} `json:"data"`
}

type Cac struct {
	State            string `json:"state,omitempty"`
	Id               int    `json:"id,omitempty"`
	Address          string `json:"address,omitempty"`
	ApprovedName     string `json:"approved_name,omitempty"`
	RcNumber         string `json:"rc_number,omitempty"`
	BranchAddress    string `json:"branch_address,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
	ClassificationId int    `json:"classification_id,omitempty"`
	Email            string `json:"email,omitempty"`
	Lga              string `json:"lga,omitempty"`
	City             string `json:"city,omitempty"`
	Status           string `json:"status,omitempty"`
}

type Shareholders []struct {
	EntityType string `json:"entity_type,omitempty"`
	Name       string `json:"name,omitempty"`
	Role       string `json:"role,omitempty"`
	Details    struct {
		Id             int    `json:"id,omitempty"`
		Address        string `json:"address,omitempty"`
		City           string `json:"city,omitempty"`
		Residence      string `json:"residence,omitempty"`
		Gender         string `json:"gender,omitempty"`
		Nationality    string `json:"nationality,omitempty"`
		Occupation     string `json:"occupation,omitempty"`
		Shares         int    `json:"shares,omitempty"`
		Status         string `json:"status,omitempty"`
		Email          string `json:"email,omitempty"`
		Phone          string `json:"phone,omitempty"`
		Dob            string `json:"dob,omitempty"`
		IdentityType   string `json:"identity_type,omitempty"`
		IdentityNumber string `json:"identity_number,omitempty"`
	} `json:"details"`
}
type VerifyAcct struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		AccountNumber string `json:"account_number,omitempty"`
		AccountName   string `json:"account_name,omitempty"`
	} `json:"data"`
}
