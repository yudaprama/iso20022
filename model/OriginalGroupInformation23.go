package model

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation23 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the group cancellation request.
	//
	// Usage: The group cancellation request identification can be used for reconciliation or to link tasks related to the cancellation request.
	GroupCancellationIdentification *Max35Text `xml:"GrpCxlId,omitempty"`

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case2 `xml:"Case,omitempty"`

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Specifies the original message name identifier to which the message refers.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Number of individual transactions contained in the original message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs,omitempty"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Indicates whether the cancellation request applies to a whole group of transactions or to individual transactions within an original group.
	GroupCancellation *GroupCancellationIndicator `xml:"GrpCxl,omitempty"`

	// Set of elements used to provide detailed information on the cancellation reason.
	CancellationReasonInformation []*CancellationReasonInformation3 `xml:"CxlRsnInf,omitempty"`
}

func (o *OriginalGroupInformation23) SetGroupCancellationIdentification(value string) {
	o.GroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation23) AddCase() *Case2 {
	o.Case = new(Case2)
	return o.Case
}

func (o *OriginalGroupInformation23) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation23) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation23) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation23) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation23) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation23) SetGroupCancellation(value string) {
	o.GroupCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalGroupInformation23) AddCancellationReasonInformation() *CancellationReasonInformation3 {
	newValue := new(CancellationReasonInformation3)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}
