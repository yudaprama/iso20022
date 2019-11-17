package model

// Choice of format for the cash settlement system.
type CashSettlementSystem4Choice struct {

	// Cash settlement system expressed as an ISO 20022 code.
	Code *CashSettlementSystem2Code `xml:"Cd"`

	// Cash settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (c *CashSettlementSystem4Choice) SetCode(value string) {
	c.Code = (*CashSettlementSystem2Code)(&value)
}

func (c *CashSettlementSystem4Choice) AddProprietary() *GenericIdentification30 {
	c.Proprietary = new(GenericIdentification30)
	return c.Proprietary
}
