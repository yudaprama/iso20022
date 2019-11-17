package model

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type Cheque9 struct {

	// Unique and unambiguous identifier for a cheque as assigned by the financial institution.
	Number *Max35Text `xml:"Nb,omitempty"`

	// Party to which a cheque is made payable.
	PayeeIdentification *PartyIdentification113 `xml:"PyeeId"`

	// Financial institution on which a cheque is drawn, that is, the financial institution that services the account of the entity that issued the cheque.
	DraweeIdentification *FinancialInstitutionIdentification10 `xml:"DrweeId,omitempty"`

	// Account owner that issues a cheque ordering the drawee bank to pay a specific amount, upon demand, to the payee.
	DrawerIdentification *PartyIdentification113 `xml:"DrwrId,omitempty"`
}

func (c *Cheque9) SetNumber(value string) {
	c.Number = (*Max35Text)(&value)
}

func (c *Cheque9) AddPayeeIdentification() *PartyIdentification113 {
	c.PayeeIdentification = new(PartyIdentification113)
	return c.PayeeIdentification
}

func (c *Cheque9) AddDraweeIdentification() *FinancialInstitutionIdentification10 {
	c.DraweeIdentification = new(FinancialInstitutionIdentification10)
	return c.DraweeIdentification
}

func (c *Cheque9) AddDrawerIdentification() *PartyIdentification113 {
	c.DrawerIdentification = new(PartyIdentification113)
	return c.DrawerIdentification
}
