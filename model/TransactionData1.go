package model

// Card transaction information to be transferred.
type TransactionData1 struct {

	// Brand name of the card.
	CardBrand *Max35Text `xml:"CardBrnd,omitempty"`

	// Card data associated with the card performing the transaction.
	CardData *PlainCardData3 `xml:"CardData,omitempty"`

	// Point of interaction (POI) performing the transaction.
	PointOfInteraction *PointOfInteraction1 `xml:"PtOfIntractn,omitempty"`

	// Details of the transaction.
	TransactionDetails *CardPaymentTransactionDetails8 `xml:"TxDtls,omitempty"`

	// PrePaid Account for funds transfer or loading.
	PrePaidAccount *CashAccount24 `xml:"PrePdAcct,omitempty"`
}

func (t *TransactionData1) SetCardBrand(value string) {
	t.CardBrand = (*Max35Text)(&value)
}

func (t *TransactionData1) AddCardData() *PlainCardData3 {
	t.CardData = new(PlainCardData3)
	return t.CardData
}

func (t *TransactionData1) AddPointOfInteraction() *PointOfInteraction1 {
	t.PointOfInteraction = new(PointOfInteraction1)
	return t.PointOfInteraction
}

func (t *TransactionData1) AddTransactionDetails() *CardPaymentTransactionDetails8 {
	t.TransactionDetails = new(CardPaymentTransactionDetails8)
	return t.TransactionDetails
}

func (t *TransactionData1) AddPrePaidAccount() *CashAccount24 {
	t.PrePaidAccount = new(CashAccount24)
	return t.PrePaidAccount
}
