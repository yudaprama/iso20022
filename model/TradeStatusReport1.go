package model

// Details of the trade status report.
type TradeStatusReport1 struct {

	// Information concerning the original message to which the TradeStatusReport is sent in response.
	OriginalMessageDetails *OriginalMessage1 `xml:"OrgnlMsgDtls"`

	// Specifies the processing status of the original message.
	Status *UndertakingStatus1Code `xml:"Sts"`

	// Set of elements used to provide detailed information on the status reason.
	StatusReason []*StatusReasonInformation8 `xml:"StsRsn,omitempty"`

	// Additional information related to the report.
	AdditionalInformation *Max35Text `xml:"AddtlInf,omitempty"`
}

func (t *TradeStatusReport1) AddOriginalMessageDetails() *OriginalMessage1 {
	t.OriginalMessageDetails = new(OriginalMessage1)
	return t.OriginalMessageDetails
}

func (t *TradeStatusReport1) SetStatus(value string) {
	t.Status = (*UndertakingStatus1Code)(&value)
}

func (t *TradeStatusReport1) AddStatusReason() *StatusReasonInformation8 {
	newValue := new(StatusReasonInformation8)
	t.StatusReason = append(t.StatusReason, newValue)
	return newValue
}

func (t *TradeStatusReport1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max35Text)(&value)
}
