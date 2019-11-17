package model

// Totals for the value of the holdings reported in the statement or page.
type TotalValueInPageAndStatement2 struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *AmountAndDirection6 `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *AmountAndDirection6 `xml:"TtlHldgsValOfStmt"`

	// Total book value of positions reported in this statement (a statement may comprise one or more messages).
	TotalBookValueOfStatement *AmountAndDirection6 `xml:"TtlBookValOfStmt,omitempty"`
}

func (t *TotalValueInPageAndStatement2) AddTotalHoldingsValueOfPage() *AmountAndDirection6 {
	t.TotalHoldingsValueOfPage = new(AmountAndDirection6)
	return t.TotalHoldingsValueOfPage
}

func (t *TotalValueInPageAndStatement2) AddTotalHoldingsValueOfStatement() *AmountAndDirection6 {
	t.TotalHoldingsValueOfStatement = new(AmountAndDirection6)
	return t.TotalHoldingsValueOfStatement
}

func (t *TotalValueInPageAndStatement2) AddTotalBookValueOfStatement() *AmountAndDirection6 {
	t.TotalBookValueOfStatement = new(AmountAndDirection6)
	return t.TotalBookValueOfStatement
}
