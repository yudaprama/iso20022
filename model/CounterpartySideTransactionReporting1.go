package model

// This is regulatory transaction reporting information from the counterparty side party.
type CounterpartySideTransactionReporting1 struct {

	// Specifies the supervisory party to which the trade needs to be reported.
	ReportingJurisdiction *Max35Text `xml:"RptgJursdctn,omitempty"`

	// Identifies the party that is responsible for reporting the trade to the trade repository.
	ReportingParty *PartyIdentification73Choice `xml:"RptgPty,omitempty"`

	// Specifies the unique transaction identifier (UTI) to be created at the time a transaction is first executed, shared with all registered entities and counterparties involved in the transaction, and used to track that particular transaction over its life. This identifier can also be known as the Unique Swap Identifier (USI). This is the UTI from the Counterparty Side party.
	CounterpartySideUniqueTransactionIdentifier []*UniqueTransactionIdentifier2 `xml:"CtrPtySdUnqTxIdr,omitempty"`
}

func (c *CounterpartySideTransactionReporting1) SetReportingJurisdiction(value string) {
	c.ReportingJurisdiction = (*Max35Text)(&value)
}

func (c *CounterpartySideTransactionReporting1) AddReportingParty() *PartyIdentification73Choice {
	c.ReportingParty = new(PartyIdentification73Choice)
	return c.ReportingParty
}

func (c *CounterpartySideTransactionReporting1) AddCounterpartySideUniqueTransactionIdentifier() *UniqueTransactionIdentifier2 {
	newValue := new(UniqueTransactionIdentifier2)
	c.CounterpartySideUniqueTransactionIdentifier = append(c.CounterpartySideUniqueTransactionIdentifier, newValue)
	return newValue
}
