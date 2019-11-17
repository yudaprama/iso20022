package model

// Choice of format for the trading capacity of the party.
type TradeOriginator4Choice struct {

	// Trading party capacity expressed as an ISO 20022 code.
	Code *OriginatorRole2Code `xml:"Cd"`

	// Trading party capacity expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TradeOriginator4Choice) SetCode(value string) {
	t.Code = (*OriginatorRole2Code)(&value)
}

func (t *TradeOriginator4Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
