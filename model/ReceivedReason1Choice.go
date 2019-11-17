package model

// Choice of formats for the reason of a received status.
type ReceivedReason1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the received status.
	Reason *ReceivedReason2Choice `xml:"Rsn"`
}

func (r *ReceivedReason1Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *ReceivedReason1Choice) AddReason() *ReceivedReason2Choice {
	r.Reason = new(ReceivedReason2Choice)
	return r.Reason
}
