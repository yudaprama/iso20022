package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties8 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount39 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount39 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`
}

func (c *CashParties8) AddDebtor() *PartyIdentificationAndAccount39 {
	c.Debtor = new(PartyIdentificationAndAccount39)
	return c.Debtor
}

func (c *CashParties8) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties8) AddCreditor() *PartyIdentificationAndAccount39 {
	c.Creditor = new(PartyIdentificationAndAccount39)
	return c.Creditor
}

func (c *CashParties8) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}
