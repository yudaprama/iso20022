package model

// Provides details on the original transactions, to which the status report message refers.
type OriginalPaymentInstruction2 struct {

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

	// Provides detailed information on the reversal reason.
	ReversalReasonInformation []*PaymentReversalReason7 `xml:"RvslRsnInf,omitempty"`

	// Provides information on the original transactions to which the reversal message refers.
	TransactionInformation []*PaymentTransaction35 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInstruction2) SetReversalPaymentInformationIdentification(value string) {
	o.ReversalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction2) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction2) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction2) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction2) SetBatchBooking(value string) {
	o.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (o *OriginalPaymentInstruction2) SetPaymentInformationReversal(value string) {
	o.PaymentInformationReversal = (*TrueFalseIndicator)(&value)
}

func (o *OriginalPaymentInstruction2) AddReversalReasonInformation() *PaymentReversalReason7 {
	newValue := new(PaymentReversalReason7)
	o.ReversalReasonInformation = append(o.ReversalReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction2) AddTransactionInformation() *PaymentTransaction35 {
	newValue := new(PaymentTransaction35)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}
