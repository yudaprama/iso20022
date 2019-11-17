package model

// Reason for the pending settlement status.
type PendingSettlementStatusChoice struct {

	// Reason for a pending status in the report.
	Reason *PendingSettlementStatusReason1 `xml:"Rsn"`

	// Proprietary identification for a reason of a specific status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`
}

func (p *PendingSettlementStatusChoice) AddReason() *PendingSettlementStatusReason1 {
	p.Reason = new(PendingSettlementStatusReason1)
	return p.Reason
}

func (p *PendingSettlementStatusChoice) AddDataSourceScheme() *GenericIdentification1 {
	p.DataSourceScheme = new(GenericIdentification1)
	return p.DataSourceScheme
}

func (p *PendingSettlementStatusChoice) SetNoReason(value string) {
	p.NoReason = (*NoReasonCode)(&value)
}
