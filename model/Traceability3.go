package model

// Identification of partners involved in exchange from the merchant to the issuer, with the relative timestamp of their exchanges.
type Traceability3 struct {

	// Identification of a partner of a message exchange.
	RelayIdentification *GenericIdentification74 `xml:"RlayId"`

	// Date and time of incoming data exchange for relaying or processing.
	TraceDateTimeIn *ISODateTime `xml:"TracDtTmIn"`

	// Date and time of the outgoing exchange for relaying or processing.
	TraceDateTimeOut *ISODateTime `xml:"TracDtTmOut"`
}

func (t *Traceability3) AddRelayIdentification() *GenericIdentification74 {
	t.RelayIdentification = new(GenericIdentification74)
	return t.RelayIdentification
}

func (t *Traceability3) SetTraceDateTimeIn(value string) {
	t.TraceDateTimeIn = (*ISODateTime)(&value)
}

func (t *Traceability3) SetTraceDateTimeOut(value string) {
	t.TraceDateTimeOut = (*ISODateTime)(&value)
}
