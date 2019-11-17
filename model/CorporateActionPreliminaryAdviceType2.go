package model

// Type of movement preliminary advice document.
type CorporateActionPreliminaryAdviceType2 struct {

	// Type of movement preliminary advice ie. new or replacement.
	Type *CorporateActionPreliminaryAdviceType1Code `xml:"Tp"`

	// Indicates whether the movement preliminary advice is sent after entitlement date.
	// Value is Yes (true) if sent after entitlement date and No (false) if sent before entitlement date.
	EligibilityIndicator *YesNoIndicator `xml:"ElgbltyInd,omitempty"`
}

func (c *CorporateActionPreliminaryAdviceType2) SetType(value string) {
	c.Type = (*CorporateActionPreliminaryAdviceType1Code)(&value)
}

func (c *CorporateActionPreliminaryAdviceType2) SetEligibilityIndicator(value string) {
	c.EligibilityIndicator = (*YesNoIndicator)(&value)
}
