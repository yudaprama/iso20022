package model

// Choice between formats for the type of corporate action event.
type CorporateActionEventType1CodeChoice struct {

	// Corporate action event type in a structured format.
	Structured *CorporateActionEventType1Code `xml:"Strd"`

	// Corporate action event type in free text form.
	Unstructured *Max35Text `xml:"Ustrd"`
}

func (c *CorporateActionEventType1CodeChoice) SetStructured(value string) {
	c.Structured = (*CorporateActionEventType1Code)(&value)
}

func (c *CorporateActionEventType1CodeChoice) SetUnstructured(value string) {
	c.Unstructured = (*Max35Text)(&value)
}
