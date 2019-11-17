package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus26Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason39 `xml:"Rsn"`
}

func (r *RejectionStatus26Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus26Choice) AddReason() *RejectionReason39 {
	newValue := new(RejectionReason39)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
