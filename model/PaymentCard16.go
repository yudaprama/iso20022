package model

// Card performing the withdrawal transaction.
type PaymentCard16 struct {

	// Entry mode used to obtain the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicates the occurrence of a fall-back on the card entry mode.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData13 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`
}

func (p *PaymentCard16) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentCard16) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentCard16) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard16) AddPlainCardData() *PlainCardData13 {
	p.PlainCardData = new(PlainCardData13)
	return p.PlainCardData
}

func (p *PaymentCard16) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard16) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}
