package model

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInformation4 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the cancellation request.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	PaymentCancellationIdentification *Max35Text `xml:"PmtCxlId,omitempty"`

	// Identifies the case.
	Case *Case2 `xml:"Case,omitempty"`

	// Unique and unambiguous identifier of the original payment information block, as assigned by the original sending party.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Information concerning the original group of transactions, to which the message refers.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the cancellation payment information group.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the cancellation payment information group, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether or not the cancellation applies to a whole group of transactions or to individual transactions within the original group.
	PaymentInformationCancellation *GroupCancellationIndicator `xml:"PmtInfCxl,omitempty"`

	// Detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation3 `xml:"CxlRsnInf,omitempty"`

	// Information concerning the original transactions, to which the cancellation request message refers.
	TransactionInformation []*PaymentTransactionInformation30 `xml:"TxInf,omitempty"`
}

func (o *OriginalPaymentInformation4) SetPaymentCancellationIdentification(value string) {
	o.PaymentCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation4) AddCase() *Case2 {
	o.Case = new(Case2)
	return o.Case
}

func (o *OriginalPaymentInformation4) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation4) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInformation4) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation4) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation4) SetPaymentInformationCancellation(value string) {
	o.PaymentInformationCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalPaymentInformation4) AddCancellationReasonInformation() *CancellationReasonInformation3 {
	newValue := new(CancellationReasonInformation3)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation4) AddTransactionInformation() *PaymentTransactionInformation30 {
	newValue := new(PaymentTransactionInformation30)
	o.TransactionInformation = append(o.TransactionInformation, newValue)
	return newValue
}
