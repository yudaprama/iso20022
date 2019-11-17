package model

// Set of elements used to provide detailed information on the number of transactions that are reported with a specific cancellation status.
type NumberOfCancellationsPerStatus1 struct {

	// Number of individual cancellation requests contained in the message, detailed per status.
	DetailedNumberOfTransactions *Max15NumericText `xml:"DtldNbOfTxs"`

	// Common cancellation request status for all individual cancellation requests reported.
	DetailedStatus *CancellationIndividualStatus1Code `xml:"DtldSts"`

	// Total of all individual amounts included in the message, irrespective of currencies, detailed per status.
	DetailedControlSum *DecimalNumber `xml:"DtldCtrlSum,omitempty"`
}

func (n *NumberOfCancellationsPerStatus1) SetDetailedNumberOfTransactions(value string) {
	n.DetailedNumberOfTransactions = (*Max15NumericText)(&value)
}

func (n *NumberOfCancellationsPerStatus1) SetDetailedStatus(value string) {
	n.DetailedStatus = (*CancellationIndividualStatus1Code)(&value)
}

func (n *NumberOfCancellationsPerStatus1) SetDetailedControlSum(value string) {
	n.DetailedControlSum = (*DecimalNumber)(&value)
}
