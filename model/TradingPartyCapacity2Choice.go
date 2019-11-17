package model

// Choice of format for the trading capacity.
type TradingPartyCapacity2Choice struct {

	// Trading capacity expressed as an ISO 20022 code.
	Code *TradingCapacity6Code `xml:"Cd"`

	// Trading capacity expressed as a proprietary code.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (t *TradingPartyCapacity2Choice) SetCode(value string) {
	t.Code = (*TradingCapacity6Code)(&value)
}

func (t *TradingPartyCapacity2Choice) AddProprietary() *GenericIdentification29 {
	t.Proprietary = new(GenericIdentification29)
	return t.Proprietary
}
