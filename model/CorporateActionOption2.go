package model

// Provides information on a CA option.
type CorporateActionOption2 struct {

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`
}

func (c *CorporateActionOption2) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionOption2) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}
