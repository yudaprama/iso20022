package model

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability5 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification76 `xml:"RlayId"`

	// Name of the outgoing protocol used by the node.
	ProtocolName *Max35Text `xml:"PrtcolNm,omitempty"`

	// Version of the protocol.
	ProtocolVersion *Max6Text `xml:"PrtcolVrsn,omitempty"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability5) AddRelayIdentification() *GenericIdentification76 {
	t.RelayIdentification = new(GenericIdentification76)
	return t.RelayIdentification
}

func (t *Traceability5) SetProtocolName(value string) {
	t.ProtocolName = (*Max35Text)(&value)
}

func (t *Traceability5) SetProtocolVersion(value string) {
	t.ProtocolVersion = (*Max6Text)(&value)
}

func (t *Traceability5) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability5) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}
