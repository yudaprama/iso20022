package model

// Sensible data associated with the payment card performing the transaction.
type PlainCardData13 struct {

	// Primary Account Number (PAN) of the card.
	PAN *Min8Max28NumericText `xml:"PAN,omitempty"`

	// Identify a card or a payment token inside a set of cards with the same PAN.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Track number 1 from magnetic stripe card.
	Track1 *Max140Text `xml:"Trck1,omitempty"`

	// Track number 2 without control characters (start /end and LRC) issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read.
	Track2 *Max140Text `xml:"Trck2,omitempty"`

	// Track number 3 from magnetic stripe card.
	Track3 *Max140Text `xml:"Trck3,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData13) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData13) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData13) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData13) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData13) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData13) SetTrack1(value string) {
	p.Track1 = (*Max140Text)(&value)
}

func (p *PlainCardData13) SetTrack2(value string) {
	p.Track2 = (*Max140Text)(&value)
}

func (p *PlainCardData13) SetTrack3(value string) {
	p.Track3 = (*Max140Text)(&value)
}

func (p *PlainCardData13) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}
