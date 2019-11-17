package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus34Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason28 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus34Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus34Choice) AddReason() *RejectionOrRepairReason28 {
	newValue := new(RejectionOrRepairReason28)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
