package model

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails3 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount3 `xml:"SfkpgAcct,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	CashDetails *AccountIdentification3Choice `xml:"CshDtls,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation5 `xml:"AddtlInf,omitempty"`

	// Capacity of customer placing the order. Primarily used by futures exchanges to indicate the CTI code (customer type indicator) as required by the US CFTC (Commodity Futures Trading Commission).
	PartyCapacity *TradingPartyCapacity2Choice `xml:"PtyCpcty,omitempty"`
}

func (c *ConfirmationPartyDetails3) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails3) AddSafekeepingAccount() *SecuritiesAccount3 {
	c.SafekeepingAccount = new(SecuritiesAccount3)
	return c.SafekeepingAccount
}

func (c *ConfirmationPartyDetails3) AddCashDetails() *AccountIdentification3Choice {
	c.CashDetails = new(AccountIdentification3Choice)
	return c.CashDetails
}

func (c *ConfirmationPartyDetails3) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails3) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails3) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}

func (c *ConfirmationPartyDetails3) AddPartyCapacity() *TradingPartyCapacity2Choice {
	c.PartyCapacity = new(TradingPartyCapacity2Choice)
	return c.PartyCapacity
}
