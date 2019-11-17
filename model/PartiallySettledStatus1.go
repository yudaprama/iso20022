package model

// Status is partially settled.
type PartiallySettledStatus1 struct {

	// Reason for the partially settled status.
	Reason *SettledStatusReason1Code `xml:"Rsn"`

	// Reason for the partially settled status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the partially settled status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the partially settled status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PartiallySettledStatus1) SetReason(value string) {
	p.Reason = (*SettledStatusReason1Code)(&value)
}

func (p *PartiallySettledStatus1) SetExtendedReason(value string) {
	p.ExtendedReason = (*Extended350Code)(&value)
}

func (p *PartiallySettledStatus1) AddDataSourceScheme() *GenericIdentification1 {
	p.DataSourceScheme = new(GenericIdentification1)
	return p.DataSourceScheme
}

func (p *PartiallySettledStatus1) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
