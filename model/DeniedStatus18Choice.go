package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus18Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the denied status.
	Reason []*DeniedReason13 `xml:"Rsn"`
}

func (d *DeniedStatus18Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus18Choice) AddReason() *DeniedReason13 {
	newValue := new(DeniedReason13)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
