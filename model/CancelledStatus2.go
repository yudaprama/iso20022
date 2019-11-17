package model

// Status is cancelled.
type CancelledStatus2 struct {

	// Reason for the cancelled status.
	Reason *CancelledStatusReason2Code `xml:"Rsn"`

	// Reason for the cancelled status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the cancelled status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`
}

func (c *CancelledStatus2) SetReason(value string) {
	c.Reason = (*CancelledStatusReason2Code)(&value)
}

func (c *CancelledStatus2) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *CancelledStatus2) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

func (c *CancelledStatus2) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}
