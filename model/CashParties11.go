package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties11 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount53 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount55 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount53 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount55 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties11) AddDebtor() *PartyIdentificationAndAccount53 {
	c.Debtor = new(PartyIdentificationAndAccount53)
	return c.Debtor
}

func (c *CashParties11) AddDebtorAgent() *PartyIdentificationAndAccount55 {
	c.DebtorAgent = new(PartyIdentificationAndAccount55)
	return c.DebtorAgent
}

func (c *CashParties11) AddCreditor() *PartyIdentificationAndAccount53 {
	c.Creditor = new(PartyIdentificationAndAccount53)
	return c.Creditor
}

func (c *CashParties11) AddCreditorAgent() *PartyIdentificationAndAccount55 {
	c.CreditorAgent = new(PartyIdentificationAndAccount55)
	return c.CreditorAgent
}
