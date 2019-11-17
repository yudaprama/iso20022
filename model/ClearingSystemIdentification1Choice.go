package model

// Information used to identify a cash clearing system.
type ClearingSystemIdentification1Choice struct {

	// Infrastructure through which the payment instruction is processed.
	ClearingSystemIdentification *CashClearingSystem3Code `xml:"ClrSysId"`

	// Proprietary clearing system service selected for a transaction.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *ClearingSystemIdentification1Choice) SetClearingSystemIdentification(value string) {
	c.ClearingSystemIdentification = (*CashClearingSystem3Code)(&value)
}

func (c *ClearingSystemIdentification1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
