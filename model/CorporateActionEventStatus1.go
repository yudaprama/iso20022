package model

// The process of notifying of an upcoming corporate action. It provides corporate action details including the different options.
type CorporateActionEventStatus1 struct {

	// Indicates whether the details provided about an event are complete or incomplete.
	EventCompletenessStatus *EventCompletenessStatus1Code `xml:"EvtCmpltnsSts"`

	// Indicates the status of the occurrence of an event.
	EventConfirmationStatus *EventConfirmationStatus1Code `xml:"EvtConfSts"`
}

func (c *CorporateActionEventStatus1) SetEventCompletenessStatus(value string) {
	c.EventCompletenessStatus = (*EventCompletenessStatus1Code)(&value)
}

func (c *CorporateActionEventStatus1) SetEventConfirmationStatus(value string) {
	c.EventConfirmationStatus = (*EventConfirmationStatus1Code)(&value)
}
