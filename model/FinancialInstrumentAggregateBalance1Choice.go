package model

// Choice of holding indicator and aggregated position of holdings.
type FinancialInstrumentAggregateBalance1Choice struct {

	// Indicates whether holdings are present.
	HoldingsIndicator *YesNoIndicator `xml:"HldgsInd"`

	// Balance of holdings.
	HoldingBalance *FinancialInstrumentAggregateBalance2 `xml:"HldgBal"`
}

func (f *FinancialInstrumentAggregateBalance1Choice) SetHoldingsIndicator(value string) {
	f.HoldingsIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAggregateBalance1Choice) AddHoldingBalance() *FinancialInstrumentAggregateBalance2 {
	f.HoldingBalance = new(FinancialInstrumentAggregateBalance2)
	return f.HoldingBalance
}
