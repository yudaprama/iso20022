package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent6Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType11Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType1Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent6Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType11Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType11Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent6Choice) AddCorporateActionEventType() *CorporateActionEventType1Choice {
	s.CorporateActionEventType = new(CorporateActionEventType1Choice)
	return s.CorporateActionEventType
}
