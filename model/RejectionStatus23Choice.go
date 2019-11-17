package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus23Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the rejection status.
	Reason []*RejectionReason36 `xml:"Rsn"`
}

func (r *RejectionStatus23Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus23Choice) AddReason() *RejectionReason36 {
	newValue := new(RejectionReason36)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
