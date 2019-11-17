package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionStatus.
	Reason []*RejectionReason10 `xml:"Rsn"`
}

func (r *RejectionStatus5Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus5Choice) AddReason() *RejectionReason10 {
	newValue := new(RejectionReason10)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
