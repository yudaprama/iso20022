package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionStatus.
	Reason []*RejectionReason6 `xml:"Rsn,omitempty"`
}

func (r *RejectionStatus3Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus3Choice) AddReason() *RejectionReason6 {
	newValue := new(RejectionReason6)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
