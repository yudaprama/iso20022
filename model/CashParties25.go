package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties25 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount111 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount112 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount111 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount112 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties25) AddDebtor() *PartyIdentificationAndAccount111 {
	c.Debtor = new(PartyIdentificationAndAccount111)
	return c.Debtor
}

func (c *CashParties25) AddDebtorAgent() *PartyIdentificationAndAccount112 {
	c.DebtorAgent = new(PartyIdentificationAndAccount112)
	return c.DebtorAgent
}

func (c *CashParties25) AddCreditor() *PartyIdentificationAndAccount111 {
	c.Creditor = new(PartyIdentificationAndAccount111)
	return c.Creditor
}

func (c *CashParties25) AddCreditorAgent() *PartyIdentificationAndAccount112 {
	c.CreditorAgent = new(PartyIdentificationAndAccount112)
	return c.CreditorAgent
}
