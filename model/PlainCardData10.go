package model

// Sensible data associated with the payment card performing the transaction.
type PlainCardData10 struct {

	// Primary Account Number (PAN) of the card, or surrogate of the PAN by a payment token.
	// It correspond to the ISO 8583 field number 2.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card or a payment token inside a set of cards with the same PAN or token.
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date from which the card can be used, expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	EffectiveDate *Max10Text `xml:"FctvDt,omitempty"`

	// Expiry date of the card or the payment token expressed either in the YYYY-MM format, or in the YYYY-MM-DD format.
	ExpiryDate *Max10Text `xml:"XpryDt,omitempty"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Track issued from the magnetic stripe card or from the ICC if the magnetic stripe was not read. The track value might be provided by a payment token.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`

	// Name of the cardholder stored on the card.
	CardholderName *Max45Text `xml:"CrdhldrNm,omitempty"`
}

func (p *PlainCardData10) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData10) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData10) SetEffectiveDate(value string) {
	p.EffectiveDate = (*Max10Text)(&value)
}

func (p *PlainCardData10) SetExpiryDate(value string) {
	p.ExpiryDate = (*Max10Text)(&value)
}

func (p *PlainCardData10) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData10) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

func (p *PlainCardData10) SetCardholderName(value string) {
	p.CardholderName = (*Max45Text)(&value)
}
