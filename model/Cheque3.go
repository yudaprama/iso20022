package model

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque3 struct {

	// Unique and unambiguous identifier for a cheque as assigned by the financial institution.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Party to which a cheque is made payable.
	PayeeIdentification *PartyIdentification2Choice `xml:"PyeeId"`

	// Financial institution on which a cheque is drawn, ie, the financial institution that services the account of the entity that issued the cheque.
	DraweeIdentification *FinancialInstitutionIdentification3Choice `xml:"DrweeId,omitempty"`

	// Account owner that issues a cheque ordering the drawee bank to pay a specific amount, upon demand, to the payee.
	DrawerIdentification *PartyIdentification2Choice `xml:"DrwrId,omitempty"`
}

func (c *Cheque3) SetNumber(value string) {
	c.Number = (*Max35Text)(&value)
}

func (c *Cheque3) AddPayeeIdentification() *PartyIdentification2Choice {
	c.PayeeIdentification = new(PartyIdentification2Choice)
	return c.PayeeIdentification
}

func (c *Cheque3) AddDraweeIdentification() *FinancialInstitutionIdentification3Choice {
	c.DraweeIdentification = new(FinancialInstitutionIdentification3Choice)
	return c.DraweeIdentification
}

func (c *Cheque3) AddDrawerIdentification() *PartyIdentification2Choice {
	c.DrawerIdentification = new(PartyIdentification2Choice)
	return c.DrawerIdentification
}
