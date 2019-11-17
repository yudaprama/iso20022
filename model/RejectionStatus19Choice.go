package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus19Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason28 `xml:"Rsn"`
}

func (r *RejectionStatus19Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus19Choice) AddReason() *RejectionReason28 {
	newValue := new(RejectionReason28)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
