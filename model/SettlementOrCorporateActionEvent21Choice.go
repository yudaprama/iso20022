package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent21Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType34Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType71Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent21Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType34Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType34Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent21Choice) AddCorporateActionEventType() *CorporateActionEventType71Choice {
	s.CorporateActionEventType = new(CorporateActionEventType71Choice)
	return s.CorporateActionEventType
}
