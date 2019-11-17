package model

// Reason for the reversed status.
type ReversedStatus2Choice struct {

	// Reason for the reversal status.
	Reason *Max350Text `xml:"Rsn"`

	// Proprietary identification of the reason for the reversed status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (r *ReversedStatus2Choice) SetReason(value string) {
	r.Reason = (*Max350Text)(&value)
}

func (r *ReversedStatus2Choice) AddDataSourceScheme() *GenericIdentification1 {
	r.DataSourceScheme = new(GenericIdentification1)
	return r.DataSourceScheme
}

func (r *ReversedStatus2Choice) SetNoSpecifiedReason(value string) {
	r.NoSpecifiedReason = (*NoReasonCode)(&value)
}
