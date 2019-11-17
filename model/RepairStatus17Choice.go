package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus17Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the repair status.
	Reason []*RepairReason13 `xml:"Rsn"`
}

func (r *RepairStatus17Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus17Choice) AddReason() *RepairReason13 {
	newValue := new(RepairReason13)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
