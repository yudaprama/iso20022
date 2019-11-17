package model

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails2 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`

	// Capacity of customer placing the order. Primarily used by futures exchanges to indicate the CTI code (customer type indicator) as required by the US CFTC (Commodity Futures Trading Commission).
	InvestorCapacity *InvestorCapacity3Choice `xml:"InvstrCpcty,omitempty"`

	// Capacity of customer placing the order. Primarily used by futures exchanges to indicate the CTI code (customer type indicator) as required by the US CFTC (Commodity Futures Trading Commission).
	TradingPartyCapacity *TradingPartyCapacity1Choice `xml:"TradgPtyCpcty,omitempty"`
}

func (c *ConfirmationPartyDetails2) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails2) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails2) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails2) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

func (c *ConfirmationPartyDetails2) AddInvestorCapacity() *InvestorCapacity3Choice {
	c.InvestorCapacity = new(InvestorCapacity3Choice)
	return c.InvestorCapacity
}

func (c *ConfirmationPartyDetails2) AddTradingPartyCapacity() *TradingPartyCapacity1Choice {
	c.TradingPartyCapacity = new(TradingPartyCapacity1Choice)
	return c.TradingPartyCapacity
}
