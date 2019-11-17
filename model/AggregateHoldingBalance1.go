package model

// Overall holding position, in a single financial instrument, held in a securities account at a specified place of safekeeping.
type AggregateHoldingBalance1 struct {

	// Identification of the financial instrument for which the balance information is specified.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Form of ownership of the holding.
	HoldingForm *FormOfSecurity1Code `xml:"HldgForm,omitempty"`

	// Specifies whether the holding is physically delivered or is a book entry only.
	HoldingPhysicalType *PhysicalTransferType1Code `xml:"HldgPhysTp,omitempty"`

	// Balance breakdown on the net position of the financial instrument.
	BalanceForFinancialInstrument []*FinancialInstrumentAggregateBalance1 `xml:"BalForFinInstrm"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (a *AggregateHoldingBalance1) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	a.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return a.FinancialInstrumentIdentification
}

func (a *AggregateHoldingBalance1) SetHoldingForm(value string) {
	a.HoldingForm = (*FormOfSecurity1Code)(&value)
}

func (a *AggregateHoldingBalance1) SetHoldingPhysicalType(value string) {
	a.HoldingPhysicalType = (*PhysicalTransferType1Code)(&value)
}

func (a *AggregateHoldingBalance1) AddBalanceForFinancialInstrument() *FinancialInstrumentAggregateBalance1 {
	newValue := new(FinancialInstrumentAggregateBalance1)
	a.BalanceForFinancialInstrument = append(a.BalanceForFinancialInstrument, newValue)
	return newValue
}

func (a *AggregateHoldingBalance1) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	a.SupplementaryData = append(a.SupplementaryData, newValue)
	return newValue
}
