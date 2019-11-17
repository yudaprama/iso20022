package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent13Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType18Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType30Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent13Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType18Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType18Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent13Choice) AddCorporateActionEventType() *CorporateActionEventType30Choice {
	s.CorporateActionEventType = new(CorporateActionEventType30Choice)
	return s.CorporateActionEventType
}
