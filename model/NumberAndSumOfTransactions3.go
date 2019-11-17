package model

// Set of elements providing the total sum of entries.
type NumberAndSumOfTransactions3 struct {

	// Number of individual entries included in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Resulting amount of the netted amounts for all debit and credit entries.
	TotalNetEntryAmount *DecimalNumber `xml:"TtlNetNtryAmt,omitempty"`

	// Indicates whether the total net entry amount is a credit or a debit amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`
}

func (n *NumberAndSumOfTransactions3) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactions3) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions3) SetTotalNetEntryAmount(value string) {
	n.TotalNetEntryAmount = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions3) SetCreditDebitIndicator(value string) {
	n.CreditDebitIndicator = (*CreditDebitCode)(&value)
}
