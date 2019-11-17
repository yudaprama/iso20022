package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus12Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the repair status.
	Reason []*RepairReason8 `xml:"Rsn"`
}

func (r *RepairStatus12Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus12Choice) AddReason() *RepairReason8 {
	newValue := new(RepairReason8)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
