package model

// Choice of format for the trading capacity of the party.
type TradeOriginator1Choice struct {

	// Trading party capacity expressed as an ISO 20022 code.
	Code *OriginatorRole2Code `xml:"Cd"`

	// Trading party capacity expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TradeOriginator1Choice) SetCode(value string) {
	t.Code = (*OriginatorRole2Code)(&value)
}

func (t *TradeOriginator1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
