package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus17Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason25 `xml:"Rsn"`
}

func (r *RejectionStatus17Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus17Choice) AddReason() *RejectionReason25 {
	newValue := new(RejectionReason25)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
