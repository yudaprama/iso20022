package model

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability2 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification76 `xml:"RlayId"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability2) AddRelayIdentification() *GenericIdentification76 {
	t.RelayIdentification = new(GenericIdentification76)
	return t.RelayIdentification
}

func (t *Traceability2) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability2) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}
