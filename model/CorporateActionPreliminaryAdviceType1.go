package model

// Type of movement preliminary advice document.
type CorporateActionPreliminaryAdviceType1 struct {

	// Type of movement preliminary advice ie. new or replacement.
	Type *CorporateActionPreliminaryAdviceType1Code `xml:"Tp"`

	// Specifies the status of the details of the event.
	ProcessingStatus *CorporateActionProcessingStatus1Choice `xml:"PrcgSts"`

	// Indicates whether the movement preliminary advice is sent after entitlement date.
	// Value is Yes (true) if sent after entitlement date and No (false) if sent before entitlement date.
	EligibilityIndicator *YesNoIndicator `xml:"ElgbltyInd,omitempty"`
}

func (c *CorporateActionPreliminaryAdviceType1) SetType(value string) {
	c.Type = (*CorporateActionPreliminaryAdviceType1Code)(&value)
}

func (c *CorporateActionPreliminaryAdviceType1) AddProcessingStatus() *CorporateActionProcessingStatus1Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus1Choice)
	return c.ProcessingStatus
}

func (c *CorporateActionPreliminaryAdviceType1) SetEligibilityIndicator(value string) {
	c.EligibilityIndicator = (*YesNoIndicator)(&value)
}
