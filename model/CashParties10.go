package model

// Specifies cash parties in the framework of a corporate action event.
type CashParties10 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount52 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount54 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount52 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties10) AddCreditor() *PartyIdentificationAndAccount52 {
	c.Creditor = new(PartyIdentificationAndAccount52)
	return c.Creditor
}

func (c *CashParties10) AddCreditorAgent() *PartyIdentificationAndAccount54 {
	c.CreditorAgent = new(PartyIdentificationAndAccount54)
	return c.CreditorAgent
}

func (c *CashParties10) AddMarketClaimCounterparty() *PartyIdentificationAndAccount52 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount52)
	return c.MarketClaimCounterparty
}
