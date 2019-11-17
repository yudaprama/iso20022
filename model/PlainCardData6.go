package model

// Sensible data associated with the payment card performing the transaction.
type PlainCardData6 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Magnetic track or equivalent payment card data.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`
}

func (p *PlainCardData6) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData6) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData6) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData6) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData6) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData6) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}
