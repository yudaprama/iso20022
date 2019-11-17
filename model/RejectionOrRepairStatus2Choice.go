package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus2Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionOrRepairStatus.
	Reason []*RejectionOrRepairReason4 `xml:"Rsn,omitempty"`
}

func (r *RejectionOrRepairStatus2Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus2Choice) AddReason() *RejectionOrRepairReason4 {
	newValue := new(RejectionOrRepairReason4)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
