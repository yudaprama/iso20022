package model

// Reject of an exchange.
type AcceptorRejection4 struct {

	// Reject reason of the message.
	RejectReason *RejectReason1Code `xml:"RjctRsn"`

	// Detailed description of an error that caused the rejection for further analysis.
	ErrorReporting []*ErrorReporting1 `xml:"ErrRptg,omitempty"`

	// Original request that caused the party to reject it.
	MessageInError *Max100KBinary `xml:"MsgInErr,omitempty"`
}

func (a *AcceptorRejection4) SetRejectReason(value string) {
	a.RejectReason = (*RejectReason1Code)(&value)
}

func (a *AcceptorRejection4) AddErrorReporting() *ErrorReporting1 {
	newValue := new(ErrorReporting1)
	a.ErrorReporting = append(a.ErrorReporting, newValue)
	return newValue
}

func (a *AcceptorRejection4) SetMessageInError(value string) {
	a.MessageInError = (*Max100KBinary)(&value)
}
