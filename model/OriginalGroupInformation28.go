package model

// Set of elements used to provide information on the original group, to which the message refers.
type OriginalGroupInformation28 struct {

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

	// Specifies the status of a group of transactions.
	GroupStatus *ExternalPaymentGroupStatus1Code `xml:"GrpSts,omitempty"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation9 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupInformation28) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation28) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation28) SetGroupStatus(value string) {
	o.GroupStatus = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalGroupInformation28) AddStatusReasonInformation() *StatusReasonInformation9 {
	newValue := new(StatusReasonInformation9)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation28) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}
