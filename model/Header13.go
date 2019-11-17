package model

// Set of characteristics related to the protocol after a rejection.
type Header13 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId,omitempty"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty,omitempty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification53 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`
}

func (h *Header13) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header13) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header13) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header13) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header13) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header13) AddRecipientParty() *GenericIdentification53 {
	h.RecipientParty = new(GenericIdentification53)
	return h.RecipientParty
}

func (h *Header13) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
