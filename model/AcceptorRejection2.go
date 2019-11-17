package model

// Reject of an exchange.
type AcceptorRejection2 struct {

	// Reject reason of the request or the advice.
	RejectReason *RejectReason1Code `xml:"RjctRsn"`

	// Additional information related to the reject of the exchange.
	AdditionalInformation *Max500Text `xml:"AddtlInf,omitempty"`

	// Original request that caused the recipient party to reject it.
	MessageInError *Max100KBinary `xml:"MsgInErr,omitempty"`
}

func (a *AcceptorRejection2) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason1Code)(&value)
}

func (a *AcceptorRejection2) SetAdditionalInformation(value string) {
	a.AdditionalInformation = (*Max500Text)(&value)
}

func (a *AcceptorRejection2) SetMessageInError(value string) {
	a.MessageInError = (*Max100KBinary)(&value)
}
