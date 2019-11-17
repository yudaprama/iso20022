package model

// Provides the details on the original group, to which the message refers.
type OriginalGroupHeader5 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original group cancellation request.
	OriginalGroupCancellationIdentification *Max35Text `xml:"OrgnlGrpCxlId,omitempty"`

	// Identifies the case.
	ResolvedCase *Case3 `xml:"RslvdCase,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group cancellation request.
	GroupCancellationStatus *GroupCancellationStatus1Code `xml:"GrpCxlSts,omitempty"`

	// Provides detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReason2 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`
}

func (o *OriginalGroupHeader5) SetOriginalGroupCancellationIdentification(value string) {
	o.OriginalGroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader5) AddResolvedCase() *Case3 {
	o.ResolvedCase = new(Case3)
	return o.ResolvedCase
}

func (o *OriginalGroupHeader5) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader5) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader5) SetGroupCancellationStatus(value string) {
	o.GroupCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalGroupHeader5) AddCancellationStatusReasonInformation() *CancellationStatusReason2 {
	newValue := new(CancellationStatusReason2)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupHeader5) AddNumberOfTransactionsPerCancellationStatus() *NumberOfTransactionsPerStatus1 {
	newValue := new(NumberOfTransactionsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}
