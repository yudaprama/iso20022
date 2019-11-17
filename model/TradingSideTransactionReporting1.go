package model

// This is regulatory transaction reporting information from the Trading Side party.
type TradingSideTransactionReporting1 struct {

	// Specifies the supervisory party to which the trade needs to be reported.
	ReportingJurisdiction *Max35Text `xml:"RptgJursdctn,omitempty"`

	// Identifies the party that is responsible for reporting the trade to the trade repository.
	ReportingParty *PartyIdentification73Choice `xml:"RptgPty,omitempty"`

	// Specifies the unique transaction identifier (UTI) to be created at the time a transaction is first executed, shared with all registered entities and counterparties involved in the transaction, and used to track that particular transaction over its life. This identifier can also be known as the Unique Swap Identifier (USI). This is the UTI from the Trading Side party.
	TradingSideUniqueTransactionIdentifier []*UniqueTransactionIdentifier2 `xml:"TradgSdUnqTxIdr,omitempty"`
}

func (t *TradingSideTransactionReporting1) SetReportingJurisdiction(value string) {
	t.ReportingJurisdiction = (*Max35Text)(&value)
}

func (t *TradingSideTransactionReporting1) AddReportingParty() *PartyIdentification73Choice {
	t.ReportingParty = new(PartyIdentification73Choice)
	return t.ReportingParty
}

func (t *TradingSideTransactionReporting1) AddTradingSideUniqueTransactionIdentifier() *UniqueTransactionIdentifier2 {
	newValue := new(UniqueTransactionIdentifier2)
	t.TradingSideUniqueTransactionIdentifier = append(t.TradingSideUniqueTransactionIdentifier, newValue)
	return newValue
}
