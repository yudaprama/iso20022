package model

// Set of element providing summary information on entries.
type TotalTransactions1 struct {

	// Indicates the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions2 `xml:"TtlNtries,omitempty"`

	// Indicates the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Indicates the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Indicates the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*NumberAndSumOfTransactionsPerBankTransactionCode1 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions1) AddTotalEntries() *NumberAndSumOfTransactions2 {
	t.TotalEntries = new(NumberAndSumOfTransactions2)
	return t.TotalEntries
}

func (t *TotalTransactions1) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions1) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions1) AddTotalEntriesPerBankTransactionCode() *NumberAndSumOfTransactionsPerBankTransactionCode1 {
	newValue := new(NumberAndSumOfTransactionsPerBankTransactionCode1)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}
