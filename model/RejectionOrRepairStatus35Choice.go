package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus35Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason29 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus35Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus35Choice) AddReason() *RejectionOrRepairReason29 {
	newValue := new(RejectionOrRepairReason29)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
