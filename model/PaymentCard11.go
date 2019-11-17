package model

// Payment card performing the transaction.
type PaymentCard11 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData9 `xml:"PlainCardData,omitempty"`

	// Masked PAN to be printed the payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

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

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard11) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard11) AddPlainCardData() *PlainCardData9 {
	p.PlainCardData = new(PlainCardData9)
	return p.PlainCardData
}

func (p *PaymentCard11) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard11) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard11) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard11) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard11) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Max35Text)(&value)
}

func (p *PaymentCard11) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard11) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}

func (p *PaymentCard11) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
