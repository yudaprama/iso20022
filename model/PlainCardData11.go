package model

// Sensitive data associated with the payment card performing the transaction.
type PlainCardData11 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`
}

func (p *PlainCardData11) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData11) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData11) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData11) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData11) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}
