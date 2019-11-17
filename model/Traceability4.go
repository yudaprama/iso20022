package model

// Identification of partners involved in exchange from the ATM to the issuer, with the relative timestamp of their exchanges.
type Traceability4 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification77 `xml:"RlayId"`

	// Identification of the relay node in the path, to enable identification of several hosts in parallel.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability4) AddRelayIdentification() *GenericIdentification77 {
	t.RelayIdentification = new(GenericIdentification77)
	return t.RelayIdentification
}

func (t *Traceability4) SetSequenceNumber(value string) {
	t.SequenceNumber = (*Max35Text)(&value)
}

func (t *Traceability4) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability4) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}
