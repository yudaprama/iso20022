package model

// Payment card performing the transaction.
type PaymentCard8 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType7 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData6 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3NumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard8) AddProtectedCardData() *ContentInformationType7 {
	p.ProtectedCardData = new(ContentInformationType7)
	return p.ProtectedCardData
}

func (p *PaymentCard8) AddPlainCardData() *PlainCardData6 {
	p.PlainCardData = new(PlainCardData6)
	return p.PlainCardData
}

func (p *PaymentCard8) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard8) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3NumericText)(&value)
}

func (p *PaymentCard8) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Exact4NumericText)(&value)
}

func (p *PaymentCard8) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard8) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
