package model

// Specifies cash parties in the framework of a corporate action event.
type CashParties28 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount120 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount121 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount120 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties28) AddCreditor() *PartyIdentificationAndAccount120 {
	c.Creditor = new(PartyIdentificationAndAccount120)
	return c.Creditor
}

func (c *CashParties28) AddCreditorAgent() *PartyIdentificationAndAccount121 {
	c.CreditorAgent = new(PartyIdentificationAndAccount121)
	return c.CreditorAgent
}

func (c *CashParties28) AddMarketClaimCounterparty() *PartyIdentificationAndAccount120 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount120)
	return c.MarketClaimCounterparty
}
