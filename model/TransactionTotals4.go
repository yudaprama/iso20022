package model

// Totals of the reconciliation.
type TransactionTotals4 struct {

	// Total of credit transactions.
	TotalCredit *TransactionTotals5 `xml:"TtlCdt"`

	// Total of debit transactions.
	TotalDebit *TransactionTotals5 `xml:"TtlDbt"`

	// Additional count which may be utilised for reconciliation.
	TotalNumber *TransactionTotals6 `xml:"TtlNb,omitempty"`
}

func (t *TransactionTotals4) AddTotalCredit() *TransactionTotals5 {
	t.TotalCredit = new(TransactionTotals5)
	return t.TotalCredit
}

func (t *TransactionTotals4) AddTotalDebit() *TransactionTotals5 {
	t.TotalDebit = new(TransactionTotals5)
	return t.TotalDebit
}

func (t *TransactionTotals4) AddTotalNumber() *TransactionTotals6 {
	t.TotalNumber = new(TransactionTotals6)
	return t.TotalNumber
}
