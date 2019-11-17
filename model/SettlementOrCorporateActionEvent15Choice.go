package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent15Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType30Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType48Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent15Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType30Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType30Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent15Choice) AddCorporateActionEventType() *CorporateActionEventType48Choice {
	s.CorporateActionEventType = new(CorporateActionEventType48Choice)
	return s.CorporateActionEventType
}
