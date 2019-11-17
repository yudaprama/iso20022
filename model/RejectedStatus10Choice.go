package model

// Specifies whether the status is provided with a reason or not.
type RejectedStatus10Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectionReason19 `xml:"Rsn"`
}

func (r *RejectedStatus10Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus10Choice) AddReason() *RejectionReason19 {
	newValue := new(RejectionReason19)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
