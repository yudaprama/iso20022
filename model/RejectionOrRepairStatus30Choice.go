package model

// Specifies whether the status is provided with a reason or not.
type RejectionOrRepairStatus30Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason for the status.
	Reason []*RejectionOrRepairReason24 `xml:"Rsn"`
}

func (r *RejectionOrRepairStatus30Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionOrRepairStatus30Choice) AddReason() *RejectionOrRepairReason24 {
	newValue := new(RejectionOrRepairReason24)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
