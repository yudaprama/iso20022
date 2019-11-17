package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus11Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason10 `xml:"Rsn,omitempty"`
}

func (r *RejectionOrRepairStatus11Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus11Choice) AddReason() *RejectionOrRepairReason10 {
	newValue := new(RejectionOrRepairReason10)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
