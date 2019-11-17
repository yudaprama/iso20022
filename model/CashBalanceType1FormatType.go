package model

// Choice of formats to  express the type of cash balance.
type CashBalanceType1FormatType struct {

	// Standard code to  specify the type of cash balance.
	Code *CashBalanceType1Code `xml:"Cd"`

	// Proprietary code to  express the type of cash balance.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (c *CashBalanceType1FormatType) SetCode(value string) {
	c.Code = (*CashBalanceType1Code)(&value)
}

func (c *CashBalanceType1FormatType) AddProprietary() *GenericIdentification13 {
	c.Proprietary = new(GenericIdentification13)
	return c.Proprietary
}
