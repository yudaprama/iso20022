package model

// Specifies whether the status is provided with a reason or not.
type DeniedStatus10Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the DeniedStatus.
	Reason []*DeniedReason5 `xml:"Rsn"`
}

func (d *DeniedStatus10Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus10Choice) AddReason() *DeniedReason5 {
	newValue := new(DeniedReason5)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
