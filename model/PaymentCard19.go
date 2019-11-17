package model

// Payment card performing the transaction.
type PaymentCard19 struct {

	// Sensitive data of the card (PlainCardData1 including the envelope), encrypted with a cryptographic key.
	ProtectedCardData *ContentInformationType10 `xml:"PrtctdCardData,omitempty"`

	// Sensitive data associated with the card performing the transaction.
	PlainCardData *PlainCardData8 `xml:"PlainCardData,omitempty"`

	// Unique reference to the card, used by both merchants and acquirers to link tokenized and non-tokenized transactions associated to the same underlying card.
	PaymentAccountReference *Max70Text `xml:"PmtAcctRef,omitempty"`

	// Masked PAN to be printed on payment receipts or displayed to the cardholder. Masked digits may be absent or replaced by another character as '*'.
	MaskedPAN *Max30Text `xml:"MskdPAN,omitempty"`

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Type of card product.
	CardProductType *CardProductType1Code `xml:"CardPdctTp,omitempty"`
}

func (p *PaymentCard19) AddProtectedCardData() *ContentInformationType10 {
	p.ProtectedCardData = new(ContentInformationType10)
	return p.ProtectedCardData
}

func (p *PaymentCard19) AddPlainCardData() *PlainCardData8 {
	p.PlainCardData = new(PlainCardData8)
	return p.PlainCardData
}

func (p *PaymentCard19) SetPaymentAccountReference(value string) {
	p.PaymentAccountReference = (*Max70Text)(&value)
}

func (p *PaymentCard19) SetMaskedPAN(value string) {
	p.MaskedPAN = (*Max30Text)(&value)
}

func (p *PaymentCard19) SetCardBrand(value string) {
	p.CardBrand = (*Max35Text)(&value)
}

func (p *PaymentCard19) SetCardProductType(value string) {
	p.CardProductType = (*CardProductType1Code)(&value)
}
