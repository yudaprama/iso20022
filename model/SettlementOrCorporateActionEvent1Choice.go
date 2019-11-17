package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent1Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType3Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType1Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent1Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType3Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType3Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent1Choice) AddCorporateActionEventType() *CorporateActionEventType1Choice {
	s.CorporateActionEventType = new(CorporateActionEventType1Choice)
	return s.CorporateActionEventType
}
