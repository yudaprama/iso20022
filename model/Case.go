package model

// Information identifying a case.
type Case struct {

	// Unique id assigned by the case creator.
	Identification *Max35Text `xml:"Id"`

	// Party that created the case.
	Creator *AnyBICIdentifier `xml:"Cretr"`

	// Set to yes if the case was closed and needs to be re-opened.
	ReopenCaseIndication *YesNoIndicator `xml:"ReopCaseIndctn,omitempty"`
}

func (c *Case) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *Case) SetCreator(value string) {
	c.Creator = (*AnyBICIdentifier)(&value)
}

func (c *Case) SetReopenCaseIndication(value string) {
	c.ReopenCaseIndication = (*YesNoIndicator)(&value)
}
