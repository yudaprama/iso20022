package model

// Reason for the cancellation pending status.
type CancellationPendingStatus1 struct {

	// Reason for the cancellation pending status.
	Reason *Max350Text `xml:"Rsn"`

	// Proprietary identification of the reason for the cancellation pending status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (c *CancellationPendingStatus1) SetReason(value string) {
	c.Reason = (*Max350Text)(&value)
}

func (c *CancellationPendingStatus1) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

func (c *CancellationPendingStatus1) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}
