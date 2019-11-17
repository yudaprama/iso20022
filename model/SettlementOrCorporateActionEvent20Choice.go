package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent20Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType24Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType71Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent20Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType24Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType24Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent20Choice) AddCorporateActionEventType() *CorporateActionEventType71Choice {
	s.CorporateActionEventType = new(CorporateActionEventType71Choice)
	return s.CorporateActionEventType
}
