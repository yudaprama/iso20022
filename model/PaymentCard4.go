package model

// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
type PaymentCard4 struct {

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData1 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Exact3NumericText `xml:"CardCtryCd,omitempty"`

	// Brand name of the card.
	CardBrand *GenericIdentification1 `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard4) AddPlainCardData() *PlainCardData1 {
	p.PlainCardData = new(PlainCardData1)
	return p.PlainCardData
}

func (p *PaymentCard4) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Exact3NumericText)(&value)
}

func (p *PaymentCard4) AddCardBrand() *GenericIdentification1 {
	p.CardBrand = new(GenericIdentification1)
	return p.CardBrand
}

func (p *PaymentCard4) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
