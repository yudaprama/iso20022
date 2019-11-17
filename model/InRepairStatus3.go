package model

// Reason for the in repair status.
type InRepairStatus3 struct {

	// Reason for the in-repair status.
	Reason *Max350Text `xml:"Rsn"`

	// Proprietary identification of the reason for the in-repair status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (i *InRepairStatus3) SetReason(value string) {
	i.Reason = (*Max350Text)(&value)
}

func (i *InRepairStatus3) AddDataSourceScheme() *GenericIdentification1 {
	i.DataSourceScheme = new(GenericIdentification1)
	return i.DataSourceScheme
}

func (i *InRepairStatus3) SetNoSpecifiedReason(value string) {
	i.NoSpecifiedReason = (*NoReasonCode)(&value)
}
