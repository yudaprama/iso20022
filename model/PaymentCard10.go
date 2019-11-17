package model

// Payment card performing the transaction.
type PaymentCard10 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData8 `xml:"PlainCardData,omitempty"`

	// Masked PAN to be printed on payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`
}

func (p *PaymentCard10) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard10) AddPlainCardData() *PlainCardData8 {
	p.PlainCardData = new(PlainCardData8)
	return p.PlainCardData
}

func (p *PaymentCard10) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard10) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard10) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}
