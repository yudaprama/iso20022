package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent17Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType32Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType56Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent17Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType32Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType32Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent17Choice) AddCorporateActionEventType() *CorporateActionEventType56Choice {
	s.CorporateActionEventType = new(CorporateActionEventType56Choice)
	return s.CorporateActionEventType
}
