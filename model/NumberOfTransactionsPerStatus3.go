package model

// Set of elements used to provide detailed information on the number of transactions that are reported with a specific transaction status.
type NumberOfTransactionsPerStatus3 struct {

	// Number of individual transactions contained in the message, detailed per status.
	DetailedNumberOfTransactions *Max15NumericText `xml:"DtldNbOfTxs"`

	// Common transaction status for all individual transactions reported.
	DetailedStatus *TransactionIndividualStatus3Code `xml:"DtldSts"`

	// Total of all individual amounts included in the message, irrespective of currencies, detailed per status.
	DetailedControlSum *DecimalNumber `xml:"DtldCtrlSum,omitempty"`
}

func (n *NumberOfTransactionsPerStatus3) SetDetailedNumberOfTransactions(value string) {
	n.DetailedNumberOfTransactions = (*Max15NumericText)(&value)
}

func (n *NumberOfTransactionsPerStatus3) SetDetailedStatus(value string) {
	n.DetailedStatus = (*TransactionIndividualStatus3Code)(&value)
}

func (n *NumberOfTransactionsPerStatus3) SetDetailedControlSum(value string) {
	n.DetailedControlSum = (*DecimalNumber)(&value)
}
