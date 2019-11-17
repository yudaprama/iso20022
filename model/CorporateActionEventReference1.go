package model

// Identification of a linked corporate action event.
type CorporateActionEventReference1 struct {

	// Identification of the linked corporate action event.
	EventIdentification *CorporateActionEventReference1Choice `xml:"EvtId"`

	// Specifies when this corporate action event is to be processed relative to a linked corporate action event.
	LinkageType *ProcessingPosition1Choice `xml:"LkgTp,omitempty"`
}

func (c *CorporateActionEventReference1) AddEventIdentification() *CorporateActionEventReference1Choice {
	c.EventIdentification = new(CorporateActionEventReference1Choice)
	return c.EventIdentification
}

func (c *CorporateActionEventReference1) AddLinkageType() *ProcessingPosition1Choice {
	c.LinkageType = new(ProcessingPosition1Choice)
	return c.LinkageType
}
