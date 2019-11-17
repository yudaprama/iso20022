package model

// Payment processes required to transfer cash from the debtor to the creditor.
type CashParties18 struct {

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentificationAndAccount80 `xml:"Dbtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *PartyIdentificationAndAccount80 `xml:"DbtrAgt,omitempty"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount80 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount80 `xml:"CdtrAgt,omitempty"`

	// Financial institution through which the transaction must pass to reach the account with institution.
	Intermediary *PartyIdentificationAndAccount80 `xml:"Intrmy,omitempty"`
}

func (c *CashParties18) AddDebtor() *PartyIdentificationAndAccount80 {
	c.Debtor = new(PartyIdentificationAndAccount80)
	return c.Debtor
}

func (c *CashParties18) AddDebtorAgent() *PartyIdentificationAndAccount80 {
	c.DebtorAgent = new(PartyIdentificationAndAccount80)
	return c.DebtorAgent
}

func (c *CashParties18) AddCreditor() *PartyIdentificationAndAccount80 {
	c.Creditor = new(PartyIdentificationAndAccount80)
	return c.Creditor
}

func (c *CashParties18) AddCreditorAgent() *PartyIdentificationAndAccount80 {
	c.CreditorAgent = new(PartyIdentificationAndAccount80)
	return c.CreditorAgent
}

func (c *CashParties18) AddIntermediary() *PartyIdentificationAndAccount80 {
	c.Intermediary = new(PartyIdentificationAndAccount80)
	return c.Intermediary
}
