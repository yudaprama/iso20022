package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties9 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount48 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount48 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties9) AddDebtor() *PartyIdentificationAndAccount48 {
	c.Debtor = new(PartyIdentificationAndAccount48)
	return c.Debtor
}

func (c *CashParties9) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties9) AddCreditor() *PartyIdentificationAndAccount48 {
	c.Creditor = new(PartyIdentificationAndAccount48)
	return c.Creditor
}

func (c *CashParties9) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}
