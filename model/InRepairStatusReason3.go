package model

// Identification of the reason for the in-repair status.
type InRepairStatusReason3 struct {

	// Reason for the in repair status.
	Reason *InRepairStatusReason1Code `xml:"Rsn"`

	// Reason of the in repair status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the in-repair status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the in-repair status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InRepairStatusReason3) SetReason(value string) {
	i.Reason = (*InRepairStatusReason1Code)(&value)
}

func (i *InRepairStatusReason3) SetExtendedReason(value string) {
	i.ExtendedReason = (*Extended350Code)(&value)
}

func (i *InRepairStatusReason3) AddDataSourceScheme() *GenericIdentification1 {
	i.DataSourceScheme = new(GenericIdentification1)
	return i.DataSourceScheme
}

func (i *InRepairStatusReason3) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
