package model

// Specifies whether the status is provided with a reason or not.
type RejectedStatus17Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectionReason27 `xml:"Rsn"`
}

func (r *RejectedStatus17Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus17Choice) AddReason() *RejectionReason27 {
	newValue := new(RejectionReason27)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
