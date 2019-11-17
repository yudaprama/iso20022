package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus13Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the repair status.
	Reason []*RepairReason9 `xml:"Rsn"`
}

func (r *RepairStatus13Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus13Choice) AddReason() *RepairReason9 {
	newValue := new(RepairReason9)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
