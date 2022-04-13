package models

type RequestCard struct {
	Status  string
	Message string
	Data    struct {
		Link string
	}
}

type CardDesign struct {
	Status  string
	Message string
	Data    struct {
		Link string
		Name string
	}
}

type CardPersonalization struct {
	Status  string
	Message string
	Data    struct {
		Id string
	}
}

type CardPersonalizationUpdate struct {
	Status  string
	Message string
	Data    struct {
		Quantity int32
		Artwork  string
	}
}
type Personalizations struct {
	Status  string
	Message string
	Data    []struct {
		Id         string
		Status     string
		Name       string
		Quantity   int
		Artwork    string
		Created_at string
		Updated_at string
	}
	Meta struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}
type Personalization struct {
	Status  string
	Message string
	Data    []struct {
		Id         string
		Status     string
		Name       string
		Quantity   int
		Artwork    string
		Created_at string
		Updated_at string
	}
	Meta struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}

type VirtualCardCreate struct {
	Status  string
	Message string
	Data    struct {
		Id string
	}
}

type Card struct {
	Status  string
	Message string
	Data    []struct {
		Disposable     bool
		Status         string
		Type           string
		Currency       string
		Brand          string
		Created_at     string
		Account_holder string
		Id             string
		Name_on_card   string
		Card_number    string
		Cvv            string
		Card_pan       string
		Last_four      string
		Expiry_month   string
		Expiry_year    string
	}
	Meta struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}

type CardDetails struct {
	Status  string
	Message string
	Data    struct {
		Id              string
		Disposable      bool
		Status          string
		Type            string
		Currency        string
		Brand           string
		Name_on_card    string
		Card_number     string
		Cvv             string
		Card_pan        string
		Last_four       string
		Expiry_month    string
		Expiry_year     string
		Created_at      string
		Balance         int
		Billing_address struct {
			State       string
			Country     string
			Street      string
			Postal_code string
		}
	}
	Account_holder string
}

type CardTxn struct {
	Status  string
	Message string
	Data    []struct {
		Id        string
		Currency  string
		Amount    int
		Type      string
		Entry     string
		Date      string
		Narration string
		Business  struct {
			Id   string
			Name string
		}
	}
	Meta struct {
		Total    int
		Pages    int
		Previous string
		Next     string
	}
}

type CardSpendLimit struct {
	Status  string
	Message string
	Data    []struct {
		Channel   string
		Amount    int
		Frequency int
	}
}
