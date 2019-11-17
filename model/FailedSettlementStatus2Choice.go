package model

// Reason for the failed settlement status.
type FailedSettlementStatus2Choice struct {

	// Reason for the failed settlement status.
	Reason *Max350Text `xml:"Rsn"`

	// Proprietary identification of the reason for the failed settlement status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (f *FailedSettlementStatus2Choice) SetReason(value string) {
	f.Reason = (*Max350Text)(&value)
}

func (f *FailedSettlementStatus2Choice) AddDataSourceScheme() *GenericIdentification1 {
	f.DataSourceScheme = new(GenericIdentification1)
	return f.DataSourceScheme
}

func (f *FailedSettlementStatus2Choice) SetNoSpecifiedReason(value string) {
	f.NoSpecifiedReason = (*NoReasonCode)(&value)
}
