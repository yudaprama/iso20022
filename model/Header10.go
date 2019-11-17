package model

// Set of characteristics related to the protocol.
type Header10 struct {

	// Identifies the type of process related to the message.
	MessageFunction *MessageFunction4Code `xml:"MsgFctn"`

	// Version of the acquirer protocol specifications.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn"`

	// Unique identification of an exchange occurrence.
	ExchangeIdentification *Max3NumericText `xml:"XchgId"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Unique identification of the partner that has initiated the exchange.
	InitiatingParty *GenericIdentification53 `xml:"InitgPty"`

	// Unique identification of the partner that is the recipient of the message exchange.
	RecipientParty *GenericIdentification53 `xml:"RcptPty,omitempty"`

	// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
	Traceability []*Traceability2 `xml:"Tracblt,omitempty"`
}

func (h *Header10) SetMessageFunction(value string) {
	h.MessageFunction = (*MessageFunction4Code)(&value)
}

func (h *Header10) SetProtocolVersion(value string) {
	h.ProtocolVersion = (*Max6Text)(&value)
}

func (h *Header10) SetExchangeIdentification(value string) {
	h.ExchangeIdentification = (*Max3NumericText)(&value)
}

func (h *Header10) SetCreationDateTime(value string) {
	h.CreationDateTime = (*ISODateTime)(&value)
}

func (h *Header10) AddInitiatingParty() *GenericIdentification53 {
	h.InitiatingParty = new(GenericIdentification53)
	return h.InitiatingParty
}

func (h *Header10) AddRecipientParty() *GenericIdentification53 {
	h.RecipientParty = new(GenericIdentification53)
	return h.RecipientParty
}

func (h *Header10) AddTraceability() *Traceability2 {
	newValue := new(Traceability2)
	h.Traceability = append(h.Traceability, newValue)
	return newValue
}
