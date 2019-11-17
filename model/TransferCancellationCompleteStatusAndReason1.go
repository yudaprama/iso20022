package model

// Status of the transfer cancellation is complete.
type TransferCancellationCompleteStatusAndReason1 struct {

	// Reason for the transfer cancellation complete status.
	Reason *CancelledStatusReason1Code `xml:"Rsn"`

	// Reason for the transfer cancellation complete status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the transfer cancellation complete status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (t *TransferCancellationCompleteStatusAndReason1) SetReason(value string) {
	t.Reason = (*CancelledStatusReason1Code)(&value)
}

func (t *TransferCancellationCompleteStatusAndReason1) SetExtendedReason(value string) {
	t.ExtendedReason = (*Extended350Code)(&value)
}

func (t *TransferCancellationCompleteStatusAndReason1) AddDataSourceScheme() *GenericIdentification1 {
	t.DataSourceScheme = new(GenericIdentification1)
	return t.DataSourceScheme
}
