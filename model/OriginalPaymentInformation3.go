package model

// Set of elements used to provide reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInformation3 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OriginalPaymentInformationCancellationIdentification *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty"`

	// Identifies the resolved case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OriginalPaymentInformationIdentification *Max35Text `xml:"OrgnlPmtInfId"`

	// Set of elements used to provide information on the original messsage.
	OriginalGroupInformation *OriginalGroupInformation3 `xml:"OrgnlGrpInf,omitempty"`

	// Number of individual transactions contained in the original payment information group.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a cancellation request, related to a payment information group.
	PaymentInformationCancellationStatus *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty"`

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`

	// Set of elements used to provide information on the original transactions to which the cancellation request message refers.
	TransactionInformationAndStatus []*PaymentTransactionInformation32 `xml:"TxInfAndSts,omitempty"`
}

func (o *OriginalPaymentInformation3) SetOriginalPaymentInformationCancellationIdentification(value string) {
	o.OriginalPaymentInformationCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation3) AddResolvedCase() *Case2 {
	o.ResolvedCase = new(Case2)
	return o.ResolvedCase
}

func (o *OriginalPaymentInformation3) SetOriginalPaymentInformationIdentification(value string) {
	o.OriginalPaymentInformationIdentification = (*Max35Text)(&value)
}

func (o *OriginalPaymentInformation3) AddOriginalGroupInformation() *OriginalGroupInformation3 {
	o.OriginalGroupInformation = new(OriginalGroupInformation3)
	return o.OriginalGroupInformation
}

func (o *OriginalPaymentInformation3) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInformation3) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInformation3) SetPaymentInformationCancellationStatus(value string) {
	o.PaymentInformationCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInformation3) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalPaymentInformation3) AddNumberOfTransactionsPerCancellationStatus() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}

func (o *OriginalPaymentInformation3) AddTransactionInformationAndStatus() *PaymentTransactionInformation32 {
	newValue := new(PaymentTransactionInformation32)
	o.TransactionInformationAndStatus = append(o.TransactionInformationAndStatus, newValue)
	return newValue
}
