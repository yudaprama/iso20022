package model

// Sensible data associated with the payment card performing the transaction.
type PlainCardData1 struct {

	// Primary Account Number (PAN) of the card, or card number.
	PAN *Min8Max28NumericText `xml:"PAN"`

	// Identify a card inside a set of cards with the same card number (PAN).
	CardSequenceNumber *Min2Max3NumericText `xml:"CardSeqNb,omitempty"`

	// Date as from which the card can be used.
	EffectiveDate *ISOYearMonth `xml:"FctvDt,omitempty"`

	// Expiry date of the card.
	ExpiryDate *ISOYearMonth `xml:"XpryDt"`

	// Services attached to the card, as defined in ISO 7813.
	ServiceCode *Exact3NumericText `xml:"SvcCd,omitempty"`

	// Magnetic track or equivalent payment card data.
	TrackData []*TrackData1 `xml:"TrckData,omitempty"`

	// Card security code (CSC) associated with the card performing the transaction.
	CardSecurityCode *CardSecurityInformation1 `xml:"CardSctyCd,omitempty"`
}

func (p *PlainCardData1) SetPAN(value string) {
	p.PAN = (*Min8Max28NumericText)(&value)
}

func (p *PlainCardData1) SetCardSequenceNumber(value string) {
	p.CardSequenceNumber = (*Min2Max3NumericText)(&value)
}

func (p *PlainCardData1) SetEffectiveDate(value string) {
	p.EffectiveDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData1) SetExpiryDate(value string) {
	p.ExpiryDate = (*ISOYearMonth)(&value)
}

func (p *PlainCardData1) SetServiceCode(value string) {
	p.ServiceCode = (*Exact3NumericText)(&value)
}

func (p *PlainCardData1) AddTrackData() *TrackData1 {
	newValue := new(TrackData1)
	p.TrackData = append(p.TrackData, newValue)
	return newValue
}

func (p *PlainCardData1) AddCardSecurityCode() *CardSecurityInformation1 {
	p.CardSecurityCode = new(CardSecurityInformation1)
	return p.CardSecurityCode
}
