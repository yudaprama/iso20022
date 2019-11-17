package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent2Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType2Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType1Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent2Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType2Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType2Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent2Choice) AddCorporateActionEventType() *CorporateActionEventType1Choice {
	s.CorporateActionEventType = new(CorporateActionEventType1Choice)
	return s.CorporateActionEventType
}
