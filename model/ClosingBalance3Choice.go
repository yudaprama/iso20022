package model

// Choice of final and intermediary balances.
type ClosingBalance3Choice struct {

	// Closing balance of the financial instrument in the statement.
	FinalClosingBalance *FinancialInstrumentQuantity1 `xml:"FnlClsgBal,omitempty"`

	// Closing Balance of this page only. Must be the intermediary opening balance of the next page (part of the same statement).
	IntermediaryClosingBalance *FinancialInstrumentQuantity1 `xml:"IntrmyClsgBal,omitempty"`
}

func (c *ClosingBalance3Choice) AddFinalClosingBalance() *FinancialInstrumentQuantity1 {
	c.FinalClosingBalance = new(FinancialInstrumentQuantity1)
	return c.FinalClosingBalance
}

func (c *ClosingBalance3Choice) AddIntermediaryClosingBalance() *FinancialInstrumentQuantity1 {
	c.IntermediaryClosingBalance = new(FinancialInstrumentQuantity1)
	return c.IntermediaryClosingBalance
}
