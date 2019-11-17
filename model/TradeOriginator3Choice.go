package model

// Choice of format for the trading capacity of the party.
type TradeOriginator3Choice struct {

	// Trading party capacity expressed as an ISO 20022 code.
	Code *OriginatorRole2Code `xml:"Cd"`

	// Trading party capacity expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TradeOriginator3Choice) SetCode(value string) {
	t.Code = (*OriginatorRole2Code)(&value)
}

func (t *TradeOriginator3Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
