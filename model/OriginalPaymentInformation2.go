package model

// Set of elements used to provide information on the original transactions, to which the status report message refers.
type OriginalPaymentInformation2 struct {

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reversed payment information group.
	// Usage: The instructing party is the party sending the reversal message and not the party that sent the original instruction that is being reversed.
	ReversalPaymentInformationIdentification *Max35Text `xml:"RvslPmtInfId,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Indicates whether or not the reversal applies to the complete original payment information group or to individual transactions within that group.
	PaymentInformationReversal *TrueFalseIndicator `xml:"PmtInfRvsl,omitempty"`

	// Set of elements used to provide detailed information on the reversal reason.
	ReversalReasonInformation []*ReversalReasonInformation6 `xml:"RvslRsnInf,omitempty"`

	// Set of elements used to provide information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransactionInformation28 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInformation2) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation2) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation2) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation2) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation2) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInformation2) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInformation2) AddReversalReasonInformation() *ReversalReasonInformation6 {
	newValue := new(ReversalReasonInformation6)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation2) AddTransactionInformation() *PaymentTransactionInformation28 {
	newValue := new(PaymentTransactionInformation28)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}
