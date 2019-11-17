package model

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation24 struct {

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original group cancellation request.
	OriginalGroupCancellationIdentification *Max35Text `xml:"OrgnlGrpCxlId,omitempty"`

	// Identifies the case.
	ResolvedCase *Case2 `xml:"RslvdCase,omitempty"`

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

	// Set of elements used to provide detailed information on the cancellation status reason.
	CancellationStatusReasonInformation []*CancellationStatusReasonInformation1 `xml:"CxlStsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical cancellation status.
	NumberOfTransactionsPerCancellationStatus []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty"`
}

func (o *OriginalGroupInformation24) SetOriginalGroupCancellationIdentification(value string) {
	o.OriginalGroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation24) AddResolvedCase() *Case2 {
	o.ResolvedCase = new(Case2)
	return o.ResolvedCase
}

func (o *OriginalGroupInformation24) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation24) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation24) SetGroupCancellationStatus(value string) {
	o.GroupCancellationStatus = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalGroupInformation24) AddCancellationStatusReasonInformation() *CancellationStatusReasonInformation1 {
	newValue := new(CancellationStatusReasonInformation1)
	o.CancellationStatusReasonInformation = append(o.CancellationStatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation24) AddNumberOfTransactionsPerCancellationStatus() *NumberOfTransactionsPerStatus1 {
	newValue := new(NumberOfTransactionsPerStatus1)
	o.NumberOfTransactionsPerCancellationStatus = append(o.NumberOfTransactionsPerCancellationStatus, newValue)
	return newValue
}
