package model

type Currency struct {
	Id    string
	Value string
}

type Currencies []Currency

func BuildCurrencies() Currencies {
	euro := Currency{
		Id:    "1",
		Value: "€",
	}

	dollar := Currency{
		Id:    "2",
		Value: "$",
	}

	return Currencies{
		euro,
		dollar,
	}
}
