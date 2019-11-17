package model

// Nature or use of the account.
type CashAccountType2 struct {

	// Account type, in a coded form.
	Code *CashAccountType4Code `xml:"Cd"`

	// Nature or use of the account in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CashAccountType2) SetCode(value string) {
	c.Code = (*CashAccountType4Code)(&value)
}

func (c *CashAccountType2) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
