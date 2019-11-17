package model

// Identification of a linked corporate action event.
type CorporateActionEventReference3 struct {

	// Identification of the linked corporate action event.
	EventIdentification *CorporateActionEventReference3Choice `xml:"EvtId"`

	// Specifies when this corporate action event is to be processed relative to a linked corporate action event.
	LinkageType *ProcessingPosition7Choice `xml:"LkgTp,omitempty"`
}

func (c *CorporateActionEventReference3) AddEventIdentification() *CorporateActionEventReference3Choice {
	c.EventIdentification = new(CorporateActionEventReference3Choice)
	return c.EventIdentification
}

func (c *CorporateActionEventReference3) AddLinkageType() *ProcessingPosition7Choice {
	c.LinkageType = new(ProcessingPosition7Choice)
	return c.LinkageType
}
