package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus8Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the repair status.
	Reason []*RepairReason6 `xml:"Rsn"`
}

func (r *RepairStatus8Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus8Choice) AddReason() *RepairReason6 {
	newValue := new(RepairReason6)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
