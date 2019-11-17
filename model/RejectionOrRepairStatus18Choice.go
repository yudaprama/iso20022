package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus18Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionOrRepairStatus.
	Reason []*RejectionOrRepairReason2 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus18Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus18Choice) AddReason() *RejectionOrRepairReason2 {
	newValue := new(RejectionOrRepairReason2)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
