package model

// Payment card performing the transaction.
type PaymentCard21 struct {

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData15 `xml:"PlainCardData,omitempty"`

	// Unique reference to the card, used by both merchants and acquirers to link tokenized and non-tokenized transactions associated to the same underlying card.
	PaymentAccountReference *Max70Text `xml:"PmtAcctRef,omitempty"`

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

	// True if the card may be used abroad.
	InternationalCard *TrueFalseIndicator `xml:"IntrnlCard,omitempty"`

	// Product that can be purchased with the card.
	AllowedProduct []*Max70Text `xml:"AllwdPdct,omitempty"`

	// Options to the service provided by the card.
	ServiceOption *Max35Text `xml:"SvcOptn,omitempty"`

	// Additional card issuer specific data.
	AdditionalCardData *Max70Text `xml:"AddtlCardData,omitempty"`
}

func (p *PaymentCard21) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard21) AddPlainCardData() *PlainCardData15 {
	p.PlainCardData = new(PlainCardData15)
	return p.PlainCardData
}

func (p *PaymentCard21) SetPaymentAccountReference(value string) {
	p.PaymentAccountReference = (*Max70Text)(&value)
}

func (p *PaymentCard21) SetIssuerBIN(value string) {
	p.IssuerBIN = (*Max15NumericText)(&value)
}

func (p *PaymentCard21) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard21) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard21) SetCardProductProfile(value string) {
	p.CardProductProfile = (*Max35Text)(&value)
}

func (p *PaymentCard21) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard21) SetInternationalCard(value string) {
	p.InternationalCard = (*TrueFalseIndicator)(&value)
}

func (p *PaymentCard21) AddAllowedProduct(value string) {
	p.AllowedProduct = append(p.AllowedProduct, (*Max70Text)(&value))
}

func (p *PaymentCard21) SetServiceOption(value string) {
	p.ServiceOption = (*Max35Text)(&value)
}

func (p *PaymentCard21) SetAdditionalCardData(value string) {
	p.AdditionalCardData = (*Max70Text)(&value)
}
