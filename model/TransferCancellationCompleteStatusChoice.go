package model

// Reason for the complete status.
type TransferCancellationCompleteStatusChoice struct {

	// Reason of the complete status.
	Reason *TransferCancellationCompleteReason1 `xml:"Rsn"`

	// Proprietary identification for a reason of a complete status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (t *TransferCancellationCompleteStatusChoice) AddReason() *TransferCancellationCompleteReason1 {
	t.Reason = new(TransferCancellationCompleteReason1)
	return t.Reason
}

func (t *TransferCancellationCompleteStatusChoice) AddDataSourceScheme() *GenericIdentification1 {
	t.DataSourceScheme = new(GenericIdentification1)
	return t.DataSourceScheme
}
