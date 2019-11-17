package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent5Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType10Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType1Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent5Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType10Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType10Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent5Choice) AddCorporateActionEventType() *CorporateActionEventType1Choice {
	s.CorporateActionEventType = new(CorporateActionEventType1Choice)
	return s.CorporateActionEventType
}
