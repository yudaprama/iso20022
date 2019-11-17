package model

// Characteristics shared by all individual items included in the currency control message.
type CurrencyControlHeader2 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Number of individual items contained in the message.
	NumberOfItems *Max15NumericText `xml:"NbOfItms"`

	// Party that receives the message.
	ReceivingParty *PartyIdentification77 `xml:"RcvgPty"`

	// Agent which sends the message.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`
}

func (c *CurrencyControlHeader2) SetMessageIdentification(value string) {
	c.MessageIdentification = (*Max35Text)(&value)
}

func (c *CurrencyControlHeader2) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}

func (c *CurrencyControlHeader2) SetNumberOfItems(value string) {
	c.NumberOfItems = (*Max15NumericText)(&value)
}

func (c *CurrencyControlHeader2) AddReceivingParty() *PartyIdentification77 {
	c.ReceivingParty = new(PartyIdentification77)
	return c.ReceivingParty
}

func (c *CurrencyControlHeader2) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.RegistrationAgent
}
