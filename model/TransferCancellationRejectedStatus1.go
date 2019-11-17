package model

// Status of the transfer cancellation is rejected.
type TransferCancellationRejectedStatus1 struct {

	// Reason for the rejected status.
	Reason *CancellationRejectedReason1Code `xml:"Rsn"`

	// Reason for the rejected status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the rejected status.
	DataSourceScheme []*GenericIdentification1 `xml:"DataSrcSchme"`
}

func (t *TransferCancellationRejectedStatus1) SetReason(value string) {
	t.Reason = (*CancellationRejectedReason1Code)(&value)
}

func (t *TransferCancellationRejectedStatus1) SetExtendedReason(value string) {
	t.ExtendedReason = (*Extended350Code)(&value)
}

func (t *TransferCancellationRejectedStatus1) AddDataSourceScheme() *GenericIdentification1 {
	newValue := new(GenericIdentification1)
	t.DataSourceScheme = append(t.DataSourceScheme, newValue)
	return newValue
}
