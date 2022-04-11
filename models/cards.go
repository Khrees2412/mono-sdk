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
