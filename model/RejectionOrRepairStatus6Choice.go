package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus6Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionOrRepairStatus.
	Reason []*RejectionOrRepairReason1 `xml:"Rsn,omitempty"`
}

func (r *RejectionOrRepairStatus6Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus6Choice) AddReason() *RejectionOrRepairReason1 {
	newValue := new(RejectionOrRepairReason1)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
