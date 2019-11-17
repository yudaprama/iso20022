package model

// Choice of transaction type, corporate action event or settlement transaction.
type SettlementOrCorporateActionEvent16Choice struct {

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType24Choice `xml:"SctiesTxTp"`

	// Specifies the type of corporate event.
	CorporateActionEventType *CorporateActionEventType48Choice `xml:"CorpActnEvtTp"`
}

func (s *SettlementOrCorporateActionEvent16Choice) AddSecuritiesTransactionType() *SecuritiesTransactionType24Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType24Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementOrCorporateActionEvent16Choice) AddCorporateActionEventType() *CorporateActionEventType48Choice {
	s.CorporateActionEventType = new(CorporateActionEventType48Choice)
	return s.CorporateActionEventType
}
