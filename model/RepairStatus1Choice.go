package model

// Specifies whether the status is provided with a reason or not.
type RepairStatus1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the RepairStatus.
	Reason []*RepairReason1 `xml:"Rsn,omitempty"`
}

func (r *RepairStatus1Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (r *RepairStatus1Choice) AddReason() *RepairReason1 {
	newValue := new(RepairReason1)
	r.Reason = append(r.Reason, newValue)
	return newValue
}
