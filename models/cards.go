package models

type RequestCard struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Link string `json:"link,omitempty"`
	} `json:"data"`
}

type CardDesign struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Link string `json:"link,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"data"`
}

type CardPersonalization struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Id string `json:"id,omitempty"`
	} `json:"data"`
}

type CardPersonalizationUpdate struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Quantity int32  `json:"quantity,omitempty"`
		Artwork  string `json:"artwork,omitempty"`
	} `json:"data"`
}
type Personalizations struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Id        string `json:"id,omitempty"`
		Status    string `json:"status,omitempty"`
		Name      string `json:"name,omitempty"`
		Quantity  int    `json:"quantity,omitempty"`
		Artwork   string `json:"artwork,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	} `json:"data,omitempty"`
	Meta struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"meta"`
}
type Personalization struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Id        string `json:"id,omitempty"`
		Status    string `json:"status,omitempty"`
		Name      string `json:"name,omitempty"`
		Quantity  int    `json:"quantity,omitempty"`
		Artwork   string `json:"artwork,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
		UpdatedAt string `json:"updated_at,omitempty"`
	} `json:"data,omitempty"`
	Meta struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"meta"`
}

type VirtualCardCreate struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Id string `json:"id,omitempty"`
	} `json:"data"`
}

type Card struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Disposable    bool   `json:"disposable,omitempty"`
		Status        string `json:"status,omitempty"`
		Type          string `json:"type,omitempty"`
		Currency      string `json:"currency,omitempty"`
		Brand         string `json:"brand,omitempty"`
		CreatedAt     string `json:"created_at,omitempty"`
		AccountHolder string `json:"account_holder,omitempty"`
		Id            string `json:"id,omitempty"`
		NameOnCard    string `json:"name_on_card,omitempty"`
		CardNumber    string `json:"card_number,omitempty"`
		Cvv           string `json:"cvv,omitempty"`
		CardPan       string `json:"card_pan,omitempty"`
		LastFour      string `json:"last_four,omitempty"`
		ExpiryMonth   string `json:"expiry_month,omitempty"`
		ExpiryYear    string `json:"expiry_year,omitempty"`
	} `json:"data,omitempty"`
	Meta struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"meta"`
}

type CardDetails struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    struct {
		Id             string `json:"id,omitempty"`
		Disposable     bool   `json:"disposable,omitempty"`
		Status         string `json:"status,omitempty"`
		Type           string `json:"type,omitempty"`
		Currency       string `json:"currency,omitempty"`
		Brand          string `json:"brand,omitempty"`
		NameOnCard     string `json:"name_on_card,omitempty"`
		CardNumber     string `json:"card_number,omitempty"`
		Cvv            string `json:"cvv,omitempty"`
		CardPan        string `json:"card_pan,omitempty"`
		LastFour       string `json:"last_four,omitempty"`
		ExpiryMonth    string `json:"expiry_month,omitempty"`
		ExpiryYear     string `json:"expiry_year,omitempty"`
		CreatedAt      string `json:"created_at,omitempty"`
		Balance        int    `json:"balance,omitempty"`
		BillingAddress struct {
			State      string `json:"state,omitempty"`
			Country    string `json:"country,omitempty"`
			Street     string `json:"street,omitempty"`
			PostalCode string `json:"postal_code,omitempty"`
		} `json:"billing_address"`
	} `json:"data"`
	AccountHolder string `json:"account_holder,omitempty"`
}

type CardTxn struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Id        string `json:"id,omitempty"`
		Currency  string `json:"currency,omitempty"`
		Amount    int    `json:"amount,omitempty"`
		Type      string `json:"type,omitempty"`
		Entry     string `json:"entry,omitempty"`
		Date      string `json:"date,omitempty"`
		Narration string `json:"narration,omitempty"`
		Business  struct {
			Id   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"business"`
	} `json:"data,omitempty"`
	Meta struct {
		Total    int    `json:"total,omitempty"`
		Pages    int    `json:"pages,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
	} `json:"meta"`
}

type CardSpendLimit struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Channel   string `json:"channel,omitempty"`
		Amount    int    `json:"amount,omitempty"`
		Frequency int    `json:"frequency,omitempty"`
	} `json:"data,omitempty"`
}
