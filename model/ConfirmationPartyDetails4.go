package model

// Parties used for acting parties that apply either to the whole message or to individual sides.
type ConfirmationPartyDetails4 struct {

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
}

func (c *ConfirmationPartyDetails4) AddIdentification() *PartyIdentification32Choice {
	c.Identification = new(PartyIdentification32Choice)
	return c.Identification
}

func (c *ConfirmationPartyDetails4) AddSafekeepingAccount() *SecuritiesAccount3 {
	c.SafekeepingAccount = new(SecuritiesAccount3)
	return c.SafekeepingAccount
}

func (c *ConfirmationPartyDetails4) AddCashDetails() *AccountIdentification3Choice {
	c.CashDetails = new(AccountIdentification3Choice)
	return c.CashDetails
}

func (c *ConfirmationPartyDetails4) AddAlternateIdentification() *AlternatePartyIdentification5 {
	c.AlternateIdentification = new(AlternatePartyIdentification5)
	return c.AlternateIdentification
}

func (c *ConfirmationPartyDetails4) SetProcessingIdentification(value string) {
	c.ProcessingIdentification = (*Max35Text)(&value)
}

func (c *ConfirmationPartyDetails4) AddAdditionalInformation() *PartyTextInformation5 {
	c.AdditionalInformation = new(PartyTextInformation5)
	return c.AdditionalInformation
}
