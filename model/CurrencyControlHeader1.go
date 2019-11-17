package model

// Characteristics shared by all individual items included in the currency control message.
type CurrencyControlHeader1 struct {

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Number of individual items contained in the message.
	NumberOfItems *Max15NumericText `xml:"NbOfItms"`

	// Party that initiates the instruction.
	InitiatingParty *PartyIdentification77 `xml:"InitgPty"`

	// Agent which forwards the message.
	ForwardingAgent *BranchAndFinancialInstitutionIdentification5 `xml:"FwdgAgt,omitempty"`
}

func (c *CurrencyControlHeader1) SetMessageIdentification(value string) {
	c.MessageIdentification = (*Max35Text)(&value)
}

func (c *CurrencyControlHeader1) SetCreationDateTime(value string) {
	c.CreationDateTime = (*ISODateTime)(&value)
}

func (c *CurrencyControlHeader1) SetNumberOfItems(value string) {
	c.NumberOfItems = (*Max15NumericText)(&value)
}

func (c *CurrencyControlHeader1) AddInitiatingParty() *PartyIdentification77 {
	c.InitiatingParty = new(PartyIdentification77)
	return c.InitiatingParty
}

func (c *CurrencyControlHeader1) AddForwardingAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.ForwardingAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.ForwardingAgent
}
