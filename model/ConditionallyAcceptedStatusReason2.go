package model

// Identification of the reason for the conditionally accepted status.
type ConditionallyAcceptedStatusReason2 struct {

	// Reason for the conditionally accepted status.
	Reason *ConditionallyAcceptedStatusReason2Code `xml:"Rsn"`

	// Reason for the conditionally accepted status.
	ExtendedReason *Extended350Code `xml:"XtndedRsn"`

	// Proprietary identification of the reason for the conditionally accepted status.
	DataSourceScheme *GenericIdentification1 `xml:"DataSrcSchme"`

	// Additional information about the conditionally accepted status reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConditionallyAcceptedStatusReason2) SetReason(value string) {
	c.Reason = (*ConditionallyAcceptedStatusReason2Code)(&value)
}

func (c *ConditionallyAcceptedStatusReason2) SetExtendedReason(value string) {
	c.ExtendedReason = (*Extended350Code)(&value)
}

func (c *ConditionallyAcceptedStatusReason2) AddDataSourceScheme() *GenericIdentification1 {
	c.DataSourceScheme = new(GenericIdentification1)
	return c.DataSourceScheme
}

func (c *ConditionallyAcceptedStatusReason2) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
