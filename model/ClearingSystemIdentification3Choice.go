package model

// Specifies the clearing system identification.
type ClearingSystemIdentification3Choice struct {

	// Infrastructure through which the payment instruction is processed, as published in an external clearing system identification code list.
	Code *ExternalCashClearingSystem1Code `xml:"Cd"`

	// Clearing system identification in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *ClearingSystemIdentification3Choice) SetCode(value string) {
	c.Code = (*ExternalCashClearingSystem1Code)(&value)
}

func (c *ClearingSystemIdentification3Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
