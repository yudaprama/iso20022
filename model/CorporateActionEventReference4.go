package model

// Identification of a linked corporate action event.
type CorporateActionEventReference4 struct {

	// Identification of the linked corporate action event.
	EventIdentification *CorporateActionEventReference4Choice `xml:"EvtId"`

	// Specifies when this corporate action event is to be processed relative to a linked corporate action event.
	LinkageType *ProcessingPosition10Choice `xml:"LkgTp,omitempty"`
}

func (c *CorporateActionEventReference4) AddEventIdentification() *CorporateActionEventReference4Choice {
	c.EventIdentification = new(CorporateActionEventReference4Choice)
	return c.EventIdentification
}

func (c *CorporateActionEventReference4) AddLinkageType() *ProcessingPosition10Choice {
	c.LinkageType = new(ProcessingPosition10Choice)
	return c.LinkageType
}
