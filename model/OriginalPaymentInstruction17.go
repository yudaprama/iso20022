package model

// Provides the details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction17 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OriginalPaymentInformationCancellationIdentification *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Provides information on the original message.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a cancellation request, related to a payment information group.
	PaymentInformationCancellationStatus *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`

	// Provides information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransaction66 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInstruction17) SetOriginalPaymentInformationCancellationIdentification(value string) {
	o.OriginalPaymentInformationCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction17) AddResolvedCase() *Case3 {
	o.ResolvedCase = new(Case3)
	return o.ResolvedCase
}

func (o *OriginalPaymentInstruction17) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction17) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInstruction17) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction17) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction17) SetPaymentInformationCancellationStatus(value string) {
	o.PaymentInformationCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction17) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction17) AddNumberOfTransactionsPerCancellationStatus() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction17) AddTransactionInformationAndStatus() *PaymentTransaction66 {
	newValue := new(PaymentTransaction66)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}
