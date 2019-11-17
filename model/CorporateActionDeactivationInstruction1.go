package model

// Provides information about the deactivation.
type CorporateActionDeactivationInstruction1 struct {

	// Date and time at which the CSD must deactivate the corporate action event or the option.
	DeactivationDateAndTime *ISODateTime `xml:"DeactvtnDtAndTm"`

	// Provides information about the option, when the deactivation instruction applies at the level of a corporate action option.
	OptionDetails []*CorporateActionOption2 `xml:"OptnDtls,omitempty"`
}

func (c *CorporateActionDeactivationInstruction1) SetDeactivationDateAndTime(value string) {
	c.DeactivationDateAndTime = (*ISODateTime)(&value)
}

func (c *CorporateActionDeactivationInstruction1) AddOptionDetails() *CorporateActionOption2 {
	newValue := new(CorporateActionOption2)
	c.OptionDetails = append(c.OptionDetails, newValue)
	return newValue
}
