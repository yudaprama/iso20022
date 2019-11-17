package model

// Status is in repair.
type InRepairStatus1 struct {

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`

	// Reason of an in repair status in the report.
	Reason *InRepairStatusReason1 `xml:"Rsn"`

	// Proprietary identification for a reason of a specific status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (i *InRepairStatus1) SetNoReason(value string) {
	i.NoReason = (*NoReasonCode)(&value)
}

func (i *InRepairStatus1) AddReason() *InRepairStatusReason1 {
	i.Reason = new(InRepairStatusReason1)
	return i.Reason
}

func (i *InRepairStatus1) AddDataSourceScheme() *GenericIdentification1 {
	i.DataSourceScheme = new(GenericIdentification1)
	return i.DataSourceScheme
}
