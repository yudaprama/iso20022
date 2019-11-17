package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus25Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason18 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus25Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus25Choice) AddReason() *RejectionOrRepairReason18 {
	newValue := new(RejectionOrRepairReason18)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
