package model

// Card performing the withdrawal transaction.
type PaymentCard17 struct {

	// Entry mode that used to obtain the card data.
	CardDataEntryMode *CardDataReading1Code `xml:"CardDataNtryMd"`

	// Indicate the occurrence of a fall-back on the card entry mode.
	FallbackIndicator *CardFallback1Code `xml:"FllbckInd,omitempty"`

	// Replacement of the message element PlainCardData by a digital envelope using a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData14 `xml:"PlainCardData,omitempty"`

	// Country code assigned to the card by the card issuer.
	CardCountryCode *Max3Text `xml:"CardCtryCd,omitempty"`

	// Currency code of the card issuer (ISO 4217 numeric code).
	CardCurrencyCode *Exact3AlphaNumericText `xml:"CardCcyCd,omitempty"`

	// Balance of the captured card or epurse if available.
	RetainedCardBalance *CurrencyAndAmount `xml:"RtndCardBal,omitempty"`
}

func (p *PaymentCard17) SetCardDataEntryMode(value string) {
	p.CardDataEntryMode = (*CardDataReading1Code)(&value)
}

func (p *PaymentCard17) SetFallbackIndicator(value string) {
	p.FallbackIndicator = (*CardFallback1Code)(&value)
}

func (p *PaymentCard17) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard17) AddPlainCardData() *PlainCardData14 {
	p.PlainCardData = new(PlainCardData14)
	return p.PlainCardData
}

func (p *PaymentCard17) SetCardCountryCode(value string) {
	p.CardCountryCode = (*Max3Text)(&value)
}

func (p *PaymentCard17) SetCardCurrencyCode(value string) {
	p.CardCurrencyCode = (*Exact3AlphaNumericText)(&value)
}

func (p *PaymentCard17) SetRetainedCardBalance(value, currency string) {
	p.RetainedCardBalance = NewCurrencyAndAmount(value, currency)
}
