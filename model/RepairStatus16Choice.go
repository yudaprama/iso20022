package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus16Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the repair status.
	Reason []*RepairReason12 `xml:"Rsn"`
}

func (r *RepairStatus16Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus16Choice) AddReason() *RepairReason12 {
	newValue := new(RepairReason12)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
