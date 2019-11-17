package model

// Nature or use of the account.
type CashAccountType2Choice struct {

	// Account type, in a coded form.
	Code *ExternalCashAccountType1Code `xml:"Cd"`

	// Nature or use of the account in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CashAccountType2Choice) SetCode(value string) {
	c.Code = (*ExternalCashAccountType1Code)(&value)
}

func (c *CashAccountType2Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
