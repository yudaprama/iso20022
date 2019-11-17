package model

// Choice of format for the trade transaction condition.
type TradeTransactionCondition4Choice struct {

	// Trade conditions expressed in a coded form as published in an external list.
	Code *ExternalTradeTransactionCondition1Code `xml:"Cd"`

	// Trade conditions expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (t *TradeTransactionCondition4Choice) SetCode(value string) {
	t.Code = (*ExternalTradeTransactionCondition1Code)(&value)
}

func (t *TradeTransactionCondition4Choice) AddProprietary() *GenericIdentification38 {
	t.Proprietary = new(GenericIdentification38)
	return t.Proprietary
}
