package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionStatus.
	Reason []*RejectionReason5 `xml:"Rsn,omitempty"`
}

func (r *RejectionStatus1Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus1Choice) AddReason() *RejectionReason5 {
	newValue := new(RejectionReason5)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
