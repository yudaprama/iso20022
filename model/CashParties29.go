package model

// Specifies cash parties in the framework of a corporate action event.
type CashParties29 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount129 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount130 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount129 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties29) AddCreditor() *PartyIdentificationAndAccount129 {
	c.Creditor = new(PartyIdentificationAndAccount129)
	return c.Creditor
}

func (c *CashParties29) AddCreditorAgent() *PartyIdentificationAndAccount130 {
	c.CreditorAgent = new(PartyIdentificationAndAccount130)
	return c.CreditorAgent
}

func (c *CashParties29) AddMarketClaimCounterparty() *PartyIdentificationAndAccount129 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount129)
	return c.MarketClaimCounterparty
}
