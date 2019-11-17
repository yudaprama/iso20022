package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus10Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason9 `xml:"Rsn,omitempty"`
}

func (r *RejectionOrRepairStatus10Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus10Choice) AddReason() *RejectionOrRepairReason9 {
	newValue := new(RejectionOrRepairReason9)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
