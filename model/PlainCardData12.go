package model

// Sensible data associated with the payment card performing the transaction.
type PlainCardData12 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`
}

func (p *PlainCardData12) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData12) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData12) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}
