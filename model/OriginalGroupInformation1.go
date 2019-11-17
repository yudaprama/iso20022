package model

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation1 struct {

	// Point to point reference assigned by the original instructing party to unambiguously identify the original group of individual transactions.
	OriginalMessageIdentification *Max35Text `xml:"OrgnlMsgId"`

	// Name assigned by the sending party to unambiguously identify the file transmitted on the network.
	NetworkFileName *Max35Text `xml:"NtwkFileNm"`

	// Specifies the original message name identifier to which the message refers, eg, pacs.003.001.01 or MT103.
	OriginalMessageNameIdentification *Max35Text `xml:"OrgnlMsgNmId"`

	// Date and time at which the original message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm,omitempty"`

	// Party that sent the file for which the status has been generated.
	//
	// Usage: this field will only be used when the content of the message could not be decoded at the receiving side.
	FileOriginator *Max35Text `xml:"FileOrgtr,omitempty"`

	// Number of individual transactions contained in the original message.
	OriginalNumberOfTransactions *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty"`

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OriginalControlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty"`

	// Specifies the status of a group of transactions.
	GroupStatus *TransactionGroupStatus1Code `xml:"GrpSts,omitempty"`

	// Detailed information on the status reason.
	StatusReasonInformation []*StatusReasonInformation1 `xml:"StsRsnInf,omitempty"`

	// Detailed information on the number of transactions for each identical individual transaction status.
	NumberOfTransactionsPerStatus []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerSts,omitempty"`
}

func (o *OriginalGroupInformation1) SetOriginalMessageIdentification(value string) {
	o.OriginalMessageIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetNetworkFileName(value string) {
	o.NetworkFileName = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalMessageNameIdentification(value string) {
	o.OriginalMessageNameIdentification = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalCreationDateTime(value string) {
	o.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalGroupInformation1) SetFileOriginator(value string) {
	o.FileOriginator = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalNumberOfTransactions(value string) {
	o.OriginalNumberOfTransactions = (*Max15NumericText)(&value)
}

func (o *OriginalGroupInformation1) SetOriginalControlSum(value string) {
	o.OriginalControlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupInformation1) SetGroupStatus(value string) {
	o.GroupStatus = (*TransactionGroupStatus1Code)(&value)
}

func (o *OriginalGroupInformation1) AddStatusReasonInformation() *StatusReasonInformation1 {
	newValue := new(StatusReasonInformation1)
	o.StatusReasonInformation = append(o.StatusReasonInformation, newValue)
	return newValue
}

func (o *OriginalGroupInformation1) AddNumberOfTransactionsPerStatus() *NumberOfTransactionsPerStatus1 {
	newValue := new(NumberOfTransactionsPerStatus1)
	o.NumberOfTransactionsPerStatus = append(o.NumberOfTransactionsPerStatus, newValue)
	return newValue
}
