package model

// Choice of format for the cash settlement system.
type CashSettlementSystem1Choice struct {

	// Cash settlement system expressed as an ISO 20022 code.
	Code *CashSettlementSystem2Code `xml:"Cd"`

	// Cash settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (c *CashSettlementSystem1Choice) SetCode(value string) {
	c.Code = (*CashSettlementSystem2Code)(&value)
}

func (c *CashSettlementSystem1Choice) AddProprietary() *GenericIdentification20 {
	c.Proprietary = new(GenericIdentification20)
	return c.Proprietary
}
