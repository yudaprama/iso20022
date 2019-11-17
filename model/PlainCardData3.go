package model

// Sensible data associated with the payment card performing the transaction provided for verification in response.
type PlainCardData3 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date as from which the card can be used.
	EffectiveDate *ISOYearMonth `xml:"FctvDt,omitempty"`

	// Expiry date of the card.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`
}

func (p *PlainCardData3) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData3) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData3) SetEffectiveDate(value string) {
	p.EffectiveDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData3) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}
