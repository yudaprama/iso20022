package model

// Status is cancelled.
type CancelledStatus1 struct {

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`

	// Reason for a cancelled status in the report.
	Reason *CancelledStatusReason1 `xml:"Rsn"`

	// Proprietary identification of a reason for a cancelled status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (c *CancelledStatus1) SetNoReason(value string) {
	c.NoReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus1) AddReason() *CancelledStatusReason1 {
	c.Reason = new(CancelledStatusReason1)
	return c.Reason
}

func (c *CancelledStatus1) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}
