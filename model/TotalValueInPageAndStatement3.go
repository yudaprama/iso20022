package model

// Totals for the value of the holdings reported in the statement or page.
type TotalValueInPageAndStatement3 struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *AmountAndDirection14 `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *AmountAndDirection14 `xml:"TtlHldgsValOfStmt"`

	// Total book value of positions reported in this statement (a statement may comprise one or more messages).
	TotalBookValueOfStatement *AmountAndDirection14 `xml:"TtlBookValOfStmt,omitempty"`

	// Total value of the holdings eligible for collateral purposes reported in this statement (a statement may comprise one or more messages).
	TotalEligibleCollateralValue *AmountAndDirection14 `xml:"TtlElgblCollVal,omitempty"`
}

func (t *TotalValueInPageAndStatement3) AddTotalHoldingsValueOfPage() *AmountAndDirection14 {
	t.TotalHoldingsValueOfPage = new(AmountAndDirection14)
	return t.TotalHoldingsValueOfPage
}

func (t *TotalValueInPageAndStatement3) AddTotalHoldingsValueOfStatement() *AmountAndDirection14 {
	t.TotalHoldingsValueOfStatement = new(AmountAndDirection14)
	return t.TotalHoldingsValueOfStatement
}

func (t *TotalValueInPageAndStatement3) AddTotalBookValueOfStatement() *AmountAndDirection14 {
	t.TotalBookValueOfStatement = new(AmountAndDirection14)
	return t.TotalBookValueOfStatement
}

func (t *TotalValueInPageAndStatement3) AddTotalEligibleCollateralValue() *AmountAndDirection14 {
	t.TotalEligibleCollateralValue = new(AmountAndDirection14)
	return t.TotalEligibleCollateralValue
}
