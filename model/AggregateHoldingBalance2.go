package model

// Overall holding position, in a single financial instrument, held in a securities account at a specified place of safekeeping.
type AggregateHoldingBalance2 struct {

	// Identification of the financial instrument for which the balance information is specified.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Balance breakdown on the net position of a financial instrument.
	BalanceForFinancialInstrument []*FinancialInstrumentAggregateBalance1 `xml:"BalForFinInstrm"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateHoldingBalance2) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateHoldingBalance2) AddBalanceForFinancialInstrument() *FinancialInstrumentAggregateBalance1 {
	newValue := new(FinancialInstrumentAggregateBalance1)
	a.BalanceForFinancialInstrument = append(a.BalanceForFinancialInstrument, newValue)
	return newValue
}

func (a *AggregateHoldingBalance2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
