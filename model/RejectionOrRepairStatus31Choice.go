package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus31Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason25 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus31Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus31Choice) AddReason() *RejectionOrRepairReason25 {
	newValue := new(RejectionOrRepairReason25)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
