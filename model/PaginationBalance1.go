package model

// Balance of a financial instrument for a specific statement page.
type PaginationBalance1 struct {

	// Opening balance of the financial instrument in the statement.
	FirstOpeningBalance *FinancialInstrumentQuantity1 `xml:"FrstOpngBal,omitempty"`

	// Opening balance of this page only. It must be the interemdiary closing balance of the previous page (part of the same statement).
	IntermediaryOpeningBalance *FinancialInstrumentQuantity1 `xml:"IntrmyOpngBal,omitempty"`

	// Closing balance of the financial instrument in the statement.
	FinalClosingBalance *FinancialInstrumentQuantity1 `xml:"FnlClsgBal,omitempty"`

	// Closing Balance of this page only. Must be the interemdiary opening balance of the next page (part of the same statement).
	IntermediaryClosingBalance *FinancialInstrumentQuantity1 `xml:"IntrmyClsgBal,omitempty"`
}

func (p *PaginationBalance1) AddFirstOpeningBalance() *FinancialInstrumentQuantity1 {
	p.FirstOpeningBalance = new(FinancialInstrumentQuantity1)
	return p.FirstOpeningBalance
}

func (p *PaginationBalance1) AddIntermediaryOpeningBalance() *FinancialInstrumentQuantity1 {
	p.IntermediaryOpeningBalance = new(FinancialInstrumentQuantity1)
	return p.IntermediaryOpeningBalance
}

func (p *PaginationBalance1) AddFinalClosingBalance() *FinancialInstrumentQuantity1 {
	p.FinalClosingBalance = new(FinancialInstrumentQuantity1)
	return p.FinalClosingBalance
}

func (p *PaginationBalance1) AddIntermediaryClosingBalance() *FinancialInstrumentQuantity1 {
	p.IntermediaryClosingBalance = new(FinancialInstrumentQuantity1)
	return p.IntermediaryClosingBalance
}
