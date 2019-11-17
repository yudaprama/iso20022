package model

// Choice of opening and intermediary balances
type OpeningBalance3Choice struct {

	// Opening balance of the financial instrument in the statement.
	FirstOpeningBalance *FinancialInstrumentQuantity1 `xml:"FrstOpngBal"`

	// Opening balance of this page only. It must be the intermediary closing balance of the previous page (part of the same statement).
	IntermediaryOpeningBalance *FinancialInstrumentQuantity1 `xml:"IntrmyOpngBal"`
}

func (o *OpeningBalance3Choice) AddFirstOpeningBalance() *FinancialInstrumentQuantity1 {
	o.FirstOpeningBalance = new(FinancialInstrumentQuantity1)
	return o.FirstOpeningBalance
}

func (o *OpeningBalance3Choice) AddIntermediaryOpeningBalance() *FinancialInstrumentQuantity1 {
	o.IntermediaryOpeningBalance = new(FinancialInstrumentQuantity1)
	return o.IntermediaryOpeningBalance
}
