package model

// Specifies whether the status is provided with a reason or not.
type RejectedStatus22Choice struct {

	// Reason not specified.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the rejected status.
	Reason []*RejectionReason35 `xml:"Rsn"`
}

func (r *RejectedStatus22Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectedStatus22Choice) AddReason() *RejectionReason35 {
	newValue := new(RejectionReason35)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
