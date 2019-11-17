package model

// Specifies the status of the details of the corporate action event.
type CorporateActionProcessingStatus1Choice struct {

	// Specifies the status of the details of the event.
	EventStatus *CorporateActionEventStatus1 `xml:"EvtSts"`

	// Indicates that the message is for information only, that is processing of client's instruction will not be supported by the Account Servicer.
	ForInformationOnlyIndicator *YesNoIndicator `xml:"ForInfOnlyInd"`
}

func (c *CorporateActionProcessingStatus1Choice) AddEventStatus() *CorporateActionEventStatus1 {
	c.EventStatus = new(CorporateActionEventStatus1)
	return c.EventStatus
}

func (c *CorporateActionProcessingStatus1Choice) SetForInformationOnlyIndicator(value string) {
	c.ForInformationOnlyIndicator = (*YesNoIndicator)(&value)
}
