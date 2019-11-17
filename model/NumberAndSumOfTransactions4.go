package model

// Set of elements providing the total sum of entries.
type NumberAndSumOfTransactions4 struct {

	// Number of individual entries included in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`

	// Resulting debit or credit amount of the netted amounts for all debit and credit entries.
	TotalNetEntry *AmountAndDirection35 `xml:"TtlNetNtry,omitempty"`
}

func (n *NumberAndSumOfTransactions4) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactions4) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}

func (n *NumberAndSumOfTransactions4) AddTotalNetEntry() *AmountAndDirection35 {
	n.TotalNetEntry = new(AmountAndDirection35)
	return n.TotalNetEntry
}
