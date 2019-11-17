package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties3 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount20 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount15 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount20 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount15 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties3) AddDebtor() *PartyIdentificationAndAccount20 {
	c.Debtor = new(PartyIdentificationAndAccount20)
	return c.Debtor
}

func (c *CashParties3) AddDebtorAgent() *PartyIdentificationAndAccount15 {
	c.DebtorAgent = new(PartyIdentificationAndAccount15)
	return c.DebtorAgent
}

func (c *CashParties3) AddCreditor() *PartyIdentificationAndAccount20 {
	c.Creditor = new(PartyIdentificationAndAccount20)
	return c.Creditor
}

func (c *CashParties3) AddCreditorAgent() *PartyIdentificationAndAccount15 {
	c.CreditorAgent = new(PartyIdentificationAndAccount15)
	return c.CreditorAgent
}
