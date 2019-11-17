package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties30 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount133 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount134 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount133 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount134 `xml:"CdtrAgt,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution.
	Intermediary *PartyIdentificationAndAccount134 `xml:"Intrmy,omitempty"`
}

func (c *CashParties30) AddDebtor() *PartyIdentificationAndAccount133 {
	c.Debtor = new(PartyIdentificationAndAccount133)
	return c.Debtor
}

func (c *CashParties30) AddDebtorAgent() *PartyIdentificationAndAccount134 {
	c.DebtorAgent = new(PartyIdentificationAndAccount134)
	return c.DebtorAgent
}

func (c *CashParties30) AddCreditor() *PartyIdentificationAndAccount133 {
	c.Creditor = new(PartyIdentificationAndAccount133)
	return c.Creditor
}

func (c *CashParties30) AddCreditorAgent() *PartyIdentificationAndAccount134 {
	c.CreditorAgent = new(PartyIdentificationAndAccount134)
	return c.CreditorAgent
}

func (c *CashParties30) AddIntermediary() *PartyIdentificationAndAccount134 {
	c.Intermediary = new(PartyIdentificationAndAccount134)
	return c.Intermediary
}
