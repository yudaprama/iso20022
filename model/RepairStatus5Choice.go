package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus5Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RepairStatus.
	Reason []*RepairReason1 `xml:"Rsn"`
}

func (r *RepairStatus5Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus5Choice) AddReason() *RepairReason1 {
	newValue := new(RepairReason1)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
