package model

// Set of elements used to provide summary information on entries.
type TotalTransactions3 struct {

	// Specifies the total number and sum of debit and credit entries.
	TotalEntries *NumberAndSumOfTransactions3 `xml:"TtlNtries,omitempty"`

	// Specifies the total number and sum of credit entries.
	TotalCreditEntries *NumberAndSumOfTransactions1 `xml:"TtlCdtNtries,omitempty"`

	// Specifies the total number and sum of debit entries.
	TotalDebitEntries *NumberAndSumOfTransactions1 `xml:"TtlDbtNtries,omitempty"`

	// Specifies the total number and sum of entries per bank transaction code.
	TotalEntriesPerBankTransactionCode []*TotalsPerBankTransactionCode2 `xml:"TtlNtriesPerBkTxCd,omitempty"`
}

func (t *TotalTransactions3) AddTotalEntries() *NumberAndSumOfTransactions3 {
	t.TotalEntries = new(NumberAndSumOfTransactions3)
	return t.TotalEntries
}

func (t *TotalTransactions3) AddTotalCreditEntries() *NumberAndSumOfTransactions1 {
	t.TotalCreditEntries = new(NumberAndSumOfTransactions1)
	return t.TotalCreditEntries
}

func (t *TotalTransactions3) AddTotalDebitEntries() *NumberAndSumOfTransactions1 {
	t.TotalDebitEntries = new(NumberAndSumOfTransactions1)
	return t.TotalDebitEntries
}

func (t *TotalTransactions3) AddTotalEntriesPerBankTransactionCode() *TotalsPerBankTransactionCode2 {
	newValue := new(TotalsPerBankTransactionCode2)
	t.TotalEntriesPerBankTransactionCode = append(t.TotalEntriesPerBankTransactionCode, newValue)
	return newValue
}
