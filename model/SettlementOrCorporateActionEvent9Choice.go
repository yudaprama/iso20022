package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent9Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType10Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType15Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent9Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType10Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType10Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent9Choice) AddCorporateActionEventType() *CorporateActionEventType15Choice {
	s.CorporateActionEventType = new(CorporateActionEventType15Choice)
	return s.CorporateActionEventType
}
