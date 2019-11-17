package model

// Status is in repair status.
type InRepairStatus2Choice struct {

	// Reason for an in repair status in the report.
	Reason *InRepairStatusReason2 `xml:"Rsn"`

	// Proprietary identification for a reason of a specific status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`
}

func (i *InRepairStatus2Choice) AddReason() *InRepairStatusReason2 {
	i.Reason = new(InRepairStatusReason2)
	return i.Reason
}

func (i *InRepairStatus2Choice) AddDataSourceScheme() *GenericIdentification1 {
	i.DataSourceScheme = new(GenericIdentification1)
	return i.DataSourceScheme
}

func (i *InRepairStatus2Choice) SetNoReason(value string) {
	i.NoReason = (*NoReasonCode)(&value)
}
