package model

// Reject of an exchange.
type AcceptorRejection3 struct {

	// Reject reason of the request or the advice.
	RejectReason *RejectReason2Code `xml:"RjctRsn"`

	// Additional information related to the reject of the exchange.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`

	// Original request that caused the recipient party to reject it.
	MessageInError *Max100KBinary `xml:"MsgInErr,omitempty"`
}

func (a *AcceptorRejection3) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason2Code)(&value)
}

func (a *AcceptorRejection3) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max500Text)(&value)
}

func (a *AcceptorRejection3) SetMessageInError(value string) {
	a.MessageInError = (*Max100KBinary)(&value)
}
