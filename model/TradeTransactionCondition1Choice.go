package model

// Choice of format for the trade transaction condition.
type TradeTransactionCondition1Choice struct {

	// Trade conditions expressed as an ISO 20022 code.
	Code *TradeTransactionCondition4Code `xml:"Cd"`

	// Trade conditions expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TradeTransactionCondition1Choice) SetCode(value string) {
	t.Code = (*TradeTransactionCondition4Code)(&value)
}

func (t *TradeTransactionCondition1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
