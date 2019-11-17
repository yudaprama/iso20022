package model

// An event determined by a corporation's board of directors, that changes the existing corporate capital structure or financial condition.
type CorporateAction9 struct {

	// Type of corporate action event, in a free-text format.
	EventType *Max35Text `xml:"EvtTp"`

	// Identification of a corporate action assigned by an official central body/entity within a given market.
	EventIdentification *Max35Text `xml:"EvtId"`
}

func (c *CorporateAction9) SetEventType(value string) {
	c.EventType = (*Max35Text)(&value)
}

func (c *CorporateAction9) SetEventIdentification(value string) {
	c.EventIdentification = (*Max35Text)(&value)
}
