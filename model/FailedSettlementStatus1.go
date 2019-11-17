package model

// Reason for the failed settlement status.
type FailedSettlementStatus1 struct {

	// Reason for the failed settlement status.
	Reason *Max350Text `xml:"Rsn"`

	// Proprietary identification of the reason for the failed settlement status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (f *FailedSettlementStatus1) SetReason(value string) {
	f.Reason = (*Max350Text)(&value)
}

func (f *FailedSettlementStatus1) AddDataSourceScheme() *GenericIdentification1 {
	f.DataSourceScheme = new(GenericIdentification1)
	return f.DataSourceScheme
}

func (f *FailedSettlementStatus1) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}
