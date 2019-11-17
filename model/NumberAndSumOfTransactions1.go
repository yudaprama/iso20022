package model

// Set of elements providing the total sum of entries.
type NumberAndSumOfTransactions1 struct {

	// Number of individual entries included in the report.
	NumberOfEntries *Max15NumericText `xml:"NbOfNtries,omitempty"`

	// Total of all individual entries included in the report.
	Sum *DecimalNumber `xml:"Sum,omitempty"`
}

func (n *NumberAndSumOfTransactions1) SetNumberOfEntries(value string) {
	n.NumberOfEntries = (*Max15NumericText)(&value)
}

func (n *NumberAndSumOfTransactions1) SetSum(value string) {
	n.Sum = (*DecimalNumber)(&value)
}
