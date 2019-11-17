package model

// Reason for the rejected status.
type TransferCancellationRejectedStatus1Choice struct {

	// Reason of the rejected status.
	Reason *TransferCancellationRejectionReason1 `xml:"Rsn"`

	// Proprietary identification for a reason of a rejected status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (t *TransferCancellationRejectedStatus1Choice) AddReason() *TransferCancellationRejectionReason1 {
	t.Reason = new(TransferCancellationRejectionReason1)
	return t.Reason
}

func (t *TransferCancellationRejectedStatus1Choice) AddDataSourceScheme() *GenericIdentification1 {
	t.DataSourceScheme = new(GenericIdentification1)
	return t.DataSourceScheme
}
