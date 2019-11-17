package model

// Reason for the pending settlement status.
type PendingSettlementStatus3Choice struct {

	// Reason for the settlement pending status.
	Reason *PendingSettlementStatusReason2Code `xml:"Rsn"`

	// Reason for the settlement pending status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the pending settlement status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (p *PendingSettlementStatus3Choice) SetReason(value string) {
	p.Reason = (*PendingSettlementStatusReason2Code)(&value)
}

func (p *PendingSettlementStatus3Choice) SetExtendedReason(value string) {
	p.ExtendedReason = (*Extended350Code)(&value)
}

func (p *PendingSettlementStatus3Choice) AddDataSourceScheme() *GenericIdentification1 {
	p.DataSourceScheme = new(GenericIdentification1)
	return p.DataSourceScheme
}

func (p *PendingSettlementStatus3Choice) SetNoSpecifiedReason(value string) {
	p.NoSpecifiedReason = (*NoReasonCode)(&value)
}
