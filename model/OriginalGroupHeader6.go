package model

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader6 struct {

	// Unique identification, as assigned by the assigner, to unambiguously identify the group cancellation request.
	//
	// Usage: The group cancellation request identification can be used for reconciliation or to link tasks related to the cancellation request.
	GroupCancellationIdentification *Max35Text `xml:"GrpCxlId,omitempty"`

	// Uniquely and unambiguously identifies an exception or an investigation workflow.
	Case *Case3 `xml:"Case,omitempty"`

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

	// Provides detailed information on the cancellation reason.
	CancellationReasonInformation []*PaymentCancellationReason3 `xml:"CxlRsnInf,omitempty"`
}

func (o *OriginalGroupHeader6) SetGroupCancellationIdentification(value string) {
	o.GroupCancellationIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader6) AddCase() *Case3 {
	o.Case = new(Case3)
	return o.Case
}

func (o *OriginalGroupHeader6) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader6) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader6) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader6) SetNumberOfTransactions(value string) {
	o.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader6) SetControlSum(value string) {
	o.ControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader6) SetGroupCancellation(value string) {
	o.GroupCancellation = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalGroupHeader6) AddCancellationReasonInformation() *PaymentCancellationReason3 {
	newValue := new(PaymentCancellationReason3)
	o.CancellationReasonInformation = append(o.CancellationReasonInformation, newValue)
	return newValue
}
