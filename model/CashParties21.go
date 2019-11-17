package model

// Specifies cash parties in the framework of a corporate action event.
type CashParties21 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount52 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount101 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount52 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties21) AddCreditor() *PartyIdentificationAndAccount52 {
	c.Creditor = new(PartyIdentificationAndAccount52)
	return c.Creditor
}

func (c *CashParties21) AddCreditorAgent() *PartyIdentificationAndAccount101 {
	c.CreditorAgent = new(PartyIdentificationAndAccount101)
	return c.CreditorAgent
}

func (c *CashParties21) AddMarketClaimCounterparty() *PartyIdentificationAndAccount52 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount52)
	return c.MarketClaimCounterparty
}
