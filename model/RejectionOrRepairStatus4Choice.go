package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus4Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionOrRepairStatus.
	Reason []*RejectionOrRepairReason3 `xml:"Rsn,omitempty"`
}

func (r *RejectionOrRepairStatus4Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus4Choice) AddReason() *RejectionOrRepairReason3 {
	newValue := new(RejectionOrRepairReason3)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
