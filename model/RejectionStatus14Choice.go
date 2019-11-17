package model

// Specifies whether the status is provided with a reason or not.
type RejectionStatus14Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RejectionStatus.
	Reason []*RejectionReason18 `xml:"Rsn"`
}

func (r *RejectionStatus14Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RejectionStatus14Choice) AddReason() *RejectionReason18 {
	newValue := new(RejectionReason18)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
