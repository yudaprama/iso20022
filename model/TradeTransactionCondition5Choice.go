package model

// Choice of format for the trade transaction condition.
type TradeTransactionCondition5Choice struct {

	// Trade conditions expressed as an ISO 20022 code.
	Code *TradeTransactionCondition4Code `xml:"Cd"`

	// Trade conditions expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TradeTransactionCondition5Choice) SetCode(value string) {
	t.Code = (*TradeTransactionCondition4Code)(&value)
}

func (t *TradeTransactionCondition5Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
