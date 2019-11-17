package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent10Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType11Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType15Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent10Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType11Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType11Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent10Choice) AddCorporateActionEventType() *CorporateActionEventType15Choice {
	s.CorporateActionEventType = new(CorporateActionEventType15Choice)
	return s.CorporateActionEventType
}
