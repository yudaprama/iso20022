package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionOrRepairStatus.
	Reason []*RejectionOrRepairReason2 `xml:"Rsn,omitempty"`
}

func (r *RejectionOrRepairStatus5Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus5Choice) AddReason() *RejectionOrRepairReason2 {
	newValue := new(RejectionOrRepairReason2)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
