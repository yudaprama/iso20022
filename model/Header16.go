package model

// Set of characteristics related to the reject of a transaction.
type Header16 struct {

	// Version of the terminal management protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the file or message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification72 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the exchange.
	RecipientParty *GenericIdentification72 `xml:"RcptPty,omitempty"`
}

func (h *Header16) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header16) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header16) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header16) AddInitiatingParty() *GenericIdentification72 {
	h.InitiatingParty = new(GenericIdentification72)
	return h.InitiatingParty
}

func (h *Header16) AddRecipientParty() *GenericIdentification72 {
	h.RecipientParty = new(GenericIdentification72)
	return h.RecipientParty
}
