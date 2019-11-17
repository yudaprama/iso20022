package model

// Reason for the cancelled status.
type CancelledStatus13Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Reason for the cancelled status.
	Reason *CancelledStatusReason3Code `xml:"Rsn"`

	// Reason for the cancelled status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the cancelled status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (c *CancelledStatus13Choice) SetNoSpecifiedReason(value string) {
	c.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (c *CancelledStatus13Choice) SetReason(value string) {
	c.Reason = (*CancelledStatusReason3Code)(&value)
}

func (c *CancelledStatus13Choice) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *CancelledStatus13Choice) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}
