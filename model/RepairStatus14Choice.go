package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus14Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the repair status.
	Reason []*RepairReason10 `xml:"Rsn"`
}

func (r *RepairStatus14Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus14Choice) AddReason() *RepairReason10 {
	newValue := new(RepairReason10)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
