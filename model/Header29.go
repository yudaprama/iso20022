package model

// Set of characteristics related to the reject of a transaction.
type Header29 struct {

	// Version of the terminal management protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Number `xml:"XchgId,omitempty"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification72 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification93 `xml:"RcptPty,omitempty"`
}

func (h *Header29) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header29) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Number)(&value)
}

func (h *Header29) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header29) AddInitiatingParty() *GenericIdentification72 {
	h.InitiatingParty = new(GenericIdentification72)
	return h.InitiatingParty
}

func (h *Header29) AddRecipientParty() *GenericIdentification93 {
	h.RecipientParty = new(GenericIdentification93)
	return h.RecipientParty
}
