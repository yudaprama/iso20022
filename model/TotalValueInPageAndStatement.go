package model

// Value of total holdings reported.
type TotalValueInPageAndStatement struct {

	// Total value of positions reported in this message.
	TotalHoldingsValueOfPage *ActiveCurrencyAndAmount `xml:"TtlHldgsValOfPg,omitempty"`

	// Total value of positions reported in this statement (a statement may comprise one or more messages).
	TotalHoldingsValueOfStatement *ActiveCurrencyAndAmount `xml:"TtlHldgsValOfStmt"`
}

func (t *TotalValueInPageAndStatement) SetTotalHoldingsValueOfPage(value, currency string) {
	t.TotalHoldingsValueOfPage = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TotalValueInPageAndStatement) SetTotalHoldingsValueOfStatement(value, currency string) {
	t.TotalHoldingsValueOfStatement = NewActiveCurrencyAndAmount(value, currency)
}
