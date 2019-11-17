package model

// Payment card performing the transaction.
type PaymentCard5 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType5 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData1 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard5) AddProtectedCardData() *ContentInformationType5 {
	p.ProtectedCardData = new(ContentInformationType5)
	return p.ProtectedCardData
}

func (p *PaymentCard5) AddPlainCardData() *PlainCardData1 {
	p.PlainCardData = new(PlainCardData1)
	return p.PlainCardData
}

func (p *PaymentCard5) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard5) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard5) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard5) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
