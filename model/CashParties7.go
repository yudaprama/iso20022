package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties7 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount38 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount38 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties7) AddDebtor() *PartyIdentificationAndAccount38 {
	c.Debtor = new(PartyIdentificationAndAccount38)
	return c.Debtor
}

func (c *CashParties7) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties7) AddCreditor() *PartyIdentificationAndAccount38 {
	c.Creditor = new(PartyIdentificationAndAccount38)
	return c.Creditor
}

func (c *CashParties7) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}
