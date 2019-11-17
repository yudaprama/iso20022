package model

// Reason for the unmatched status.
type TransferUnmatchedStatus3Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the unmatched status.
	Reason *TransferUnmatchedReason2Code `xml:"Rsn"`

	// Reason for the unmatched status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the unmatched status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (t *TransferUnmatchedStatus3Choice) SetNoSpecifiedReason(value string) {
	t.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (t *TransferUnmatchedStatus3Choice) SetReason(value string) {
	t.Reason = (*TransferUnmatchedReason2Code)(&value)
}

func (t *TransferUnmatchedStatus3Choice) SetExtendedReason(value string) {
	t.ExtendedReason = (*Extended350Code)(&value)
}

func (t *TransferUnmatchedStatus3Choice) AddDataSourceScheme() *GenericIdentification1 {
	t.DataSourceScheme = new(GenericIdentification1)
	return t.DataSourceScheme
}
