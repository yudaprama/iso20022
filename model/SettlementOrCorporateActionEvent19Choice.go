package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent19Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType18Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType56Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent19Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType18Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType18Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent19Choice) AddCorporateActionEventType() *CorporateActionEventType56Choice {
	s.CorporateActionEventType = new(CorporateActionEventType56Choice)
	return s.CorporateActionEventType
}
