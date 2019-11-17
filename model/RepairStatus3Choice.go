package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RepairStatus.
	Reason []*RepairReason3 `xml:"Rsn,omitempty"`
}

func (r *RepairStatus3Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus3Choice) AddReason() *RepairReason3 {
	newValue := new(RepairReason3)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
