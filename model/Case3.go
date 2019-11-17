package model

// Provides further details to identify an investigation case.
type Case3 struct {

	// Uniquely identifies the case.
	Identification *Max35Text `xml:"Id"`

	// Party that created the investigation case.
	Creator *Party12Choice `xml:"Cretr"`

	// Indicates whether or not the case was previously closed and is now re-opened.
	ReopenCaseIndication *YesNoIndicator `xml:"ReopCaseIndctn,omitempty"`
}

func (c *Case3) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *Case3) AddCreator() *Party12Choice {
	c.Creator = new(Party12Choice)
	return c.Creator
}

func (c *Case3) SetReopenCaseIndication(value string) {
	c.ReopenCaseIndication = (*YesNoIndicator)(&value)
}
