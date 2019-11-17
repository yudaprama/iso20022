package model

// Payment card performing the transaction.
type PaymentCard9 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData7 `xml:"PlainCardData,omitempty"`

	// Bank identifier number of the issuer for routing purpose.
	IssuerBIN *Max15NumericText `xml:"IssrBIN,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Defines a category of cards related to the acceptance processing rules defined by the acquirer.
	CardProductProfile *Max35Text `xml:"CardPdctPrfl,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard9) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard9) AddPlainCardData() *PlainCardData7 {
	p.PlainCardData = new(PlainCardData7)
	return p.PlainCardData
}

func (p *PaymentCard9) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard9) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard9) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard9) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Max35Text)(&value)
}

func (p *PaymentCard9) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard9) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
