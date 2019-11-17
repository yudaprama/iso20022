package model

// Choice of format for the cash settlement system.
type CashSettlementSystem3Choice struct {

	// Cash settlement system expressed as an ISO 20022 code.
	Code *CashSettlementSystem2Code `xml:"Cd"`

	// Cash settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (c *CashSettlementSystem3Choice) SetCode(value string) {
	c.Code = (*CashSettlementSystem2Code)(&value)
}

func (c *CashSettlementSystem3Choice) AddProprietary() *GenericIdentification38 {
	c.Proprietary = new(GenericIdentification38)
	return c.Proprietary
}
