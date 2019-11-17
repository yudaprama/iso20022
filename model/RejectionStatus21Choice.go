package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus21Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason30 `xml:"Rsn"`
}

func (r *RejectionStatus21Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus21Choice) AddReason() *RejectionReason30 {
	newValue := new(RejectionReason30)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
