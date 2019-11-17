package model

// Balance of a financial instrument for a specific statement page.
type PaginationBalance2 struct {

	// Opening balance of the financial instrument in the statement or of the intermediary opening balance of the page of the statement.
	OpeningBalance *OpeningBalance3Choice `xml:"OpngBal,omitempty"`

	// Closing balance of the financial instrument in the statement or of the intermediary closing balance of the page of the statement
	ClosingBalance *ClosingBalance3Choice `xml:"ClsgBal,omitempty"`
}

func (p *PaginationBalance2) AddOpeningBalance() *OpeningBalance3Choice {
	p.OpeningBalance = new(OpeningBalance3Choice)
	return p.OpeningBalance
}

func (p *PaginationBalance2) AddClosingBalance() *ClosingBalance3Choice {
	p.ClosingBalance = new(ClosingBalance3Choice)
	return p.ClosingBalance
}
