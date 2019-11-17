package model

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability1 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification31 `xml:"RlayId"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability1) AddRelayIdentification() *GenericIdentification31 {
	t.RelayIdentification = new(GenericIdentification31)
	return t.RelayIdentification
}

func (t *Traceability1) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability1) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}
