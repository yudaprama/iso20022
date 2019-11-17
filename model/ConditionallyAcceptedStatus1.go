package model

// Status that is accepted under certain conditions.
type ConditionallyAcceptedStatus1 struct {

	// Indicates that there is no reason available or to report.
	NoReason *NoReasonCode `xml:"NoRsn"`

	// Reason for a conditionally accepted status in the report.
	Reason *ConditionallyAcceptedStatusReason1 `xml:"Rsn"`

	// Proprietary identification of a reason for a conditionally accepted status in the report.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`
}

func (c *ConditionallyAcceptedStatus1) SetNoReason(value string) {
	c.NoReason = (*NoReasonCode)(&value)
}

func (c *ConditionallyAcceptedStatus1) AddReason() *ConditionallyAcceptedStatusReason1 {
	c.Reason = new(ConditionallyAcceptedStatusReason1)
	return c.Reason
}

func (c *ConditionallyAcceptedStatus1) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}
