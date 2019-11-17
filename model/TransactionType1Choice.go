package model

// Choice of transaction type or corporation action event type.
type TransactionType1Choice struct {

	// Type of investment fund transaction.
	TransactionType *TransactionType2Choice `xml:"TxTp"`

	// Type of corporate action event.
	CorporateActionType *CorporateAction1Choice `xml:"CorpActnTp"`
}

func (t *TransactionType1Choice) AddTransactionType() *TransactionType2Choice {
	t.TransactionType = new(TransactionType2Choice)
	return t.TransactionType
}

func (t *TransactionType1Choice) AddCorporateActionType() *CorporateAction1Choice {
	t.CorporateActionType = new(CorporateAction1Choice)
	return t.CorporateActionType
}
