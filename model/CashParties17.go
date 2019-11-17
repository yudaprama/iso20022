package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties17 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount39 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount50 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount39 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount50 `xml:"CdtrAgt,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution.
	Intermediary *PartyIdentificationAndAccount50 `xml:"Intrmy,omitempty"`
}

func (c *CashParties17) AddDebtor() *PartyIdentificationAndAccount39 {
	c.Debtor = new(PartyIdentificationAndAccount39)
	return c.Debtor
}

func (c *CashParties17) AddDebtorAgent() *PartyIdentificationAndAccount50 {
	c.DebtorAgent = new(PartyIdentificationAndAccount50)
	return c.DebtorAgent
}

func (c *CashParties17) AddCreditor() *PartyIdentificationAndAccount39 {
	c.Creditor = new(PartyIdentificationAndAccount39)
	return c.Creditor
}

func (c *CashParties17) AddCreditorAgent() *PartyIdentificationAndAccount50 {
	c.CreditorAgent = new(PartyIdentificationAndAccount50)
	return c.CreditorAgent
}

func (c *CashParties17) AddIntermediary() *PartyIdentificationAndAccount50 {
	c.Intermediary = new(PartyIdentificationAndAccount50)
	return c.Intermediary
}
