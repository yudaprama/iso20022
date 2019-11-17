package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent14Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType19Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType30Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent14Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType19Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType19Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent14Choice) AddCorporateActionEventType() *CorporateActionEventType30Choice {
	s.CorporateActionEventType = new(CorporateActionEventType30Choice)
	return s.CorporateActionEventType
}
