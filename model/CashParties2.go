package model

// Specifies cash parties in the framework of a corporate action event.
type CashParties2 struct {

	// Party to which an amount of money is due.
	Creditor *PartyIdentificationAndAccount17 `xml:"Cdtr,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *PartyIdentificationAndAccount18 `xml:"CdtrAgt,omitempty"`

	// Party that has reimbursed the account owner with funds to which they were legally entitled.
	MarketClaimCounterparty *PartyIdentificationAndAccount17 `xml:"MktClmCtrPty,omitempty"`
}

func (c *CashParties2) AddCreditor() *PartyIdentificationAndAccount17 {
	c.Creditor = new(PartyIdentificationAndAccount17)
	return c.Creditor
}

func (c *CashParties2) AddCreditorAgent() *PartyIdentificationAndAccount18 {
	c.CreditorAgent = new(PartyIdentificationAndAccount18)
	return c.CreditorAgent
}

func (c *CashParties2) AddMarketClaimCounterparty() *PartyIdentificationAndAccount17 {
	c.MarketClaimCounterparty = new(PartyIdentificationAndAccount17)
	return c.MarketClaimCounterparty
}
