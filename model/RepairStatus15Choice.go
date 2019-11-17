package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus15Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the repair status.
	Reason []*RepairReason11 `xml:"Rsn"`
}

func (r *RepairStatus15Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus15Choice) AddReason() *RepairReason11 {
	newValue := new(RepairReason11)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
