package model

// Status is unmatched.
type TransferUnmatchedStatus struct {

	// Reason for an unmatched status in the report.
	Reason *TransferUnmatchedStatusReason1 `xml:"Rsn"`

	// Proprietary identification for a reason of a specific status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`
}

func (t *TransferUnmatchedStatus) AddReason() *TransferUnmatchedStatusReason1 {
	t.Reason = new(TransferUnmatchedStatusReason1)
	return t.Reason
}

func (t *TransferUnmatchedStatus) AddDataSourceScheme() *GenericIdentification1 {
	t.DataSourceScheme = new(GenericIdentification1)
	return t.DataSourceScheme
}

func (t *TransferUnmatchedStatus) SetNoReason(value string) {
	t.NoReason = (*NoReasonCode)(&value)
}
