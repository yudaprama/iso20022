package model

// Set of elements used to identify a case.
type Case2 struct {

	// Uniquely identifies the case.
	Identification *Max35Text `xml:"Id"`

	// Party that created the investigation case.
	Creator *Party7Choice `xml:"Cretr"`

	// Indicates whether or not the case was previously closed and is now re-opened.
	ReopenCaseIndication *YesNoIndicator `xml:"ReopCaseIndctn,omitempty"`
}

func (c *Case2) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *Case2) AddCreator() *Party7Choice {
	c.Creator = new(Party7Choice)
	return c.Creator
}

func (c *Case2) SetReopenCaseIndication(value string) {
	c.ReopenCaseIndication = (*YesNoIndicator)(&value)
}
