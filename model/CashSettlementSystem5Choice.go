package model

// Choice of format for the cash settlement system.
type CashSettlementSystem5Choice struct {

	// Cash settlement system expressed as an ISO 20022 code.
	Code *CashSettlementSystem2Code `xml:"Cd"`

	// Cash settlement system expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (c *CashSettlementSystem5Choice) SetCode(value string) {
	c.Code = (*CashSettlementSystem2Code)(&value)
}

func (c *CashSettlementSystem5Choice) AddProprietary() *GenericIdentification47 {
	c.Proprietary = new(GenericIdentification47)
	return c.Proprietary
}
