package model

// Reason for the cancellation pending status.
type CancellationPendingStatus7Choice struct {

	// Reason for the cancellation pending status.
	Reason *Max350Text `xml:"Rsn"`

	// Proprietary identification of the reason for the cancellation pending status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (c *CancellationPendingStatus7Choice) SetReason(value string) {
	c.Reason = (*Max350Text)(&value)
}

func (c *CancellationPendingStatus7Choice) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

func (c *CancellationPendingStatus7Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}
