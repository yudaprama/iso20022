package model

// Specification of the commission type.
type CommissionType1 struct {

	// Structured format.
	Structured *CommissionType5Code `xml:"Strd"`

	// Additional information about the type of commission.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CommissionType1) SetStructured(value string) {
	c.Structured = (*CommissionType5Code)(&value)
}

func (c *CommissionType1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
