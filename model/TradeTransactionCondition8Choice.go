package model

// Choice of format for the trade transaction conditions.
type TradeTransactionCondition8Choice struct {

	// Trade condition expressed as a code.
	Code *TradeTransactionCondition5Code `xml:"Cd"`

	// Trade condition expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TradeTransactionCondition8Choice) SetCode(value string) {
	t.Code = (*TradeTransactionCondition5Code)(&value)
}

func (t *TradeTransactionCondition8Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
