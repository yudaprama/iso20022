package model

// Set of elements used to provide summary information on entries.
type TotalTransactions4 struct {

	// Specifies the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions4 `xml:"TtlNtries,omitempty"`

	// Specifies the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Specifies the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*TotalsPerBankTransactionCode3 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions4) AddTotalEntries() *NumberAndSumOfTransactions4 {
	t.TotalEntries = new(NumberAndSumOfTransactions4)
	return t.TotalEntries
}

func (t *TotalTransactions4) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions4) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions4) AddTotalEntriesPerBankTransactionCode() *TotalsPerBankTransactionCode3 {
	newValue := new(TotalsPerBankTransactionCode3)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}
