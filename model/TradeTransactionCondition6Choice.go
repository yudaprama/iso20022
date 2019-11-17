package model

// Choice of format for the trade transaction condition.
type TradeTransactionCondition6Choice struct {

	// Trade conditions expressed as an ISO 20022 code.
	Code *TradeTransactionCondition4Code `xml:"Cd"`

	// Trade conditions expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TradeTransactionCondition6Choice) SetCode(value string) {
	t.Code = (*TradeTransactionCondition4Code)(&value)
}

func (t *TradeTransactionCondition6Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
