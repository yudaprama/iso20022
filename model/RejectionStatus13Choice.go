package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus13Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionStatus.
	Reason []*RejectionReason17 `xml:"Rsn"`
}

func (r *RejectionStatus13Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus13Choice) AddReason() *RejectionReason17 {
	newValue := new(RejectionReason17)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
